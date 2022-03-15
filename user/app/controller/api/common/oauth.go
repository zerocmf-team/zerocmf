/**
** @创建时间: 2021/11/28 10:51
** @作者　　: return
** @描述　　:
 */

package common

import (
	"context"
	"fmt"
	"gincmf/app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/db"
	"github.com/gincmf/bootstrap/util"
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

var (
	oauthConfig oauth2.Config
)

type Oauth struct {
	Srv   *server.Server
	Store *mysql.Store
}

func (oauth *Oauth) NewServer(tenants ...string) *Oauth {

	database := "gincmf"
	if len(tenants) > 0 {
		if tenants[0] != "" {
			database = "tenant_" + tenants[0]
		}
	}

	conf := config.Config()

	authServerURL := "http://localhost:" + strconv.Itoa(conf.App.Port)

	oauthConfig = oauth2.Config{
		ClientID:     database,
		ClientSecret: "999999",
		Scopes:       []string{"all"},
		RedirectURL:  authServerURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/authorize",
			TokenURL: authServerURL + "/token",
		},
	}

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	username := conf.Database.Username
	pwd := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port

	dsn := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlStore := mysql.NewDefaultStore(
		mysql.NewConfig(dsn),
	)

	// token memory store
	manager.MapTokenStorage(mysqlStore)

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()

	clientStore.Set(oauthConfig.ClientID, &models.Client{
		ID:     oauthConfig.ClientID,
		Secret: oauthConfig.ClientSecret,
		Domain: authServerURL,
	})

	manager.MapClientStorage(clientStore)
	srv := server.NewDefaultServer(manager)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return &Oauth{
		Srv:   srv,
		Store: mysqlStore,
	}
}

func Routes(e *gin.Engine) {
	e.POST("api/oauth/token", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		// 验证用户账号密码
		user := model.User{}
		tx := db.Db().Where("user_login = ?", username).First(&user)
		if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
			new(controller.Rest).Error(c, "查询用户失败", nil)
			return
		}

		//验证密码
		if util.GetMd5(password) != user.UserPass {
			new(controller.Rest).Error(c, "账号密码不正确！", nil)
			return
		}

		// 更新最后登录记录
		u := model.User{
			LastLoginIp: c.ClientIP(),
			LastLoginAt: time.Now().Unix(),
		}
		db.Db().Where("id = ?", user.Id).Updates(&u)

		tokenExp := "168"

		exp, err := strconv.Atoi(tokenExp)

		if err != nil {
			fmt.Println("err", err.Error())
			new(controller.Rest).Error(c, "失效时间应该是整数，单位为小时！", nil)
			return
		}

		oauth := new(Oauth).NewServer()
		defer oauth.Store.Close()
		srv := oauth.Srv

		duration := time.Duration(exp) * time.Hour

		req := &server.AuthorizeRequest{
			RedirectURI:    oauthConfig.RedirectURL,
			ResponseType:   "code",
			ClientID:       oauthConfig.ClientID,
			State:          "jwt",
			Scope:          "all",
			UserID:         strconv.Itoa(user.Id),
			AccessTokenExp: duration,
			Request:        c.Request,
		}

		ti, err := srv.GetAuthorizeToken(c, req)

		if err != nil {
			new(controller.Rest).Error(c, "系统错误："+err.Error(), nil)
			return
		}

		code := ti.GetCode()

		token, err := oauthConfig.Exchange(context.Background(), code)

		if err != nil {
			fmt.Println("err", err)
			new(controller.Rest).Error(c, "获取失败："+err.Error(), nil)
			return
		}

		new(controller.Rest).Success(c, "获取成功！", token)

	})

	e.POST("api/oauth/refresh", func(c *gin.Context) {

		refreshToken := c.Query("refresh_token")
		if refreshToken == "" {
			new(controller.Rest).Error(c, "refresh_token不能为空！", nil)
			return
		}
		token := &oauth2.Token{RefreshToken: refreshToken}
		tkr := oauthConfig.TokenSource(context.Background(), token)
		tk, err := tkr.Token()

		if err != nil {
			new(controller.Rest).Error(c, "获取失败："+err.Error(), gin.H{
				"error":             "invalid_client",
				"error_description": "Client authentication failed",
			})
			return
		}
		new(controller.Rest).Success(c, "获取成功！", tk)
	})

	//e.POST("/authorize", func(c *gin.Context) {
	//
	//	oauth := new(Oauth).NewServer()
	//	defer oauth.Store.Close()
	//	srv := oauth.Srv
	//
	//	err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
	//	if err != nil {
	//		new(controller.Rest).Error(c, err.Error(), nil)
	//	}
	//})
	//
	e.POST("/token", func(c *gin.Context) {
		oauth := new(Oauth).NewServer()
		defer oauth.Store.Close()
		srv := oauth.Srv

		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println("/token err", err.Error())
			return
		}
	})
}
