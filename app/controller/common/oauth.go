/**
** @创建时间: 2020/10/5 2:34 下午
** @作者　　: return
** @描述　　:
 */
package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/util"
	"golang.org/x/oauth2"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"net/http"
	"strconv"
	"time"
)

var (
	globalOauthToken  *oauth2.Token
	globalTenantToken *oauth2.Token
	rc                controller.RestController
	Srv               *server.Server
)

func RegisterOauthRouter() {

	conf := cmf.Conf()
	// 创建userToken表
	// db.AutoMigrate(userToken{})
	authServerURL := "http://localhost:"+conf.App.Port
	clientSecret := util.GetMd5(conf.Database.AuthCode)
	fmt.Println("clientSecret", clientSecret)

	config := oauth2.Config{
		ClientID:     "1",
		ClientSecret: clientSecret,
		Scopes:       []string{"all"},
		RedirectURL:  authServerURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/authorize",
			TokenURL: authServerURL + "/token",
		},
	}

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store

	manager.MustTokenStorage(store.NewFileTokenStore("../data.db"))

	// generate jwt access token
	accessToken := generates.NewJWTAccessGenerate([]byte("gincmf"+conf.Database.AuthCode), jwt.SigningMethodHS512)
	manager.MapAccessGenerate(accessToken)
	clientStore := store.NewClientStore()
	clientStore.Set(config.ClientID, &models.Client{
		ID:     config.ClientID,
		Secret: config.ClientSecret,
		Domain: authServerURL,
	})

	manager.MapClientStorage(clientStore)
	Srv = server.NewServer(server.NewConfig(), manager)
	Srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {

		type User struct {
			Id           int
			UserType     int
			Gender       int
			Birthday     int
			UserLogin    string `gorm:"type:varchar(60);not null"`
			UserPass     string `gorm:"type:varchar(64);not null"`
			UserNickname string
			Avatar       string
			Signature    string
			Mobile       string
		}

		u := &User{}
		userResult := cmf.Db.First(u, "user_login = ?", username) // 查询
		if userResult.RowsAffected > 0 {
			//验证密码
			if util.GetMd5(password) == u.UserPass {
				userID = strconv.Itoa(u.Id)
			}
			return userID, nil
		}
		return "", errors.New("当前用户不存在！")

	})

	cmf.Post("api/oauth/token", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		autoLogin := c.DefaultPostForm("autoLogin", "false")

		tokenExp := "24"
		if autoLogin == "true" {
			tokenExp = "168"
		}

		exp, err := strconv.Atoi(tokenExp)

		if err != nil {
			fmt.Println("err", err.Error())
			rc.Error(c, "失效时间应该是整数，单位为小时！", nil)
			return
		}

		fn := Srv.PasswordAuthorizationHandler
		userID, err := fn(username, password)

		if userID == "" {
			rc.Error(c, "账号密码不正确！", nil)
			return
		}

		//cuserId,_ := strconv.Atoi(userID)

		duration := time.Duration(exp) * time.Hour

		req := &server.AuthorizeRequest{
			RedirectURI:    config.RedirectURL,
			ResponseType:   "code",
			ClientID:       config.ClientID,
			State:          "jwt",
			Scope:          "all",
			UserID:         userID,
			AccessTokenExp: duration,
			Request:        c.Request,
		}

		ti, err := Srv.GetAuthorizeToken(req)
		if err != nil {
			fmt.Println(err.Error())
		}

		code := ti.GetCode()

		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			//panic(err.Error())
		}

		//timeNow := time.Now()
		// createAt := timeNow.Unix()
		// expireAt := createAt +  int64(exp) * 60
		globalOauthToken = token
		fmt.Println("globalToken", globalOauthToken)

		//userToken{
		//	UserId: userId,
		//	ExpireAt: expireAt,
		//	CreateAt: createAt,
		//	Token: token.AccessToken,
		//}

		c.JSON(http.StatusOK, token)
	})

	cmf.Post("api/oauth/refresh", func(c *gin.Context) {

		grantType := c.Query("grant_type")
		if grantType != "refresh_token" {
			rc.Error(c, "grant_type类型错误！", nil)
			return
		}

		refreshToken := c.Query("refresh_token")
		if refreshToken == "" {
			rc.Error(c, "refresh_token不能为空！", nil)
			return
		}
		token := &oauth2.Token{RefreshToken: refreshToken}
		tkr := config.TokenSource(context.Background(), token)
		tk, err := tkr.Token()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":             "invalid_client",
				"error_description": "Client authentication failed",
			})
			return
		}

		c.JSON(http.StatusOK, tk)
	})

	cmf.Post("/token", func(c *gin.Context) {
		fmt.Println("token request")
		err := Srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			panic(err.Error())
		}
	})

}