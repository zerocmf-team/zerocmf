/**
** @创建时间: 2022/3/29 18:47
** @作者　　: return
** @描述　　:
 */

package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/oauth2"
	"log"
	"strconv"
)

type oauth struct {
	Config oauth2.Config
	Srv    *server.Server
	Store  *mysql.Store
}

type Config struct {
	Port     int
	Database struct {
		Type     string
		Host     string
		Database string
		Username string
		Password string
		Port     int
		Charset  string
		Prefix   string
		AuthCode string
	}
}

func NewConf(conf Config, tenantId string) (oauthConfig oauth2.Config) {
	database := conf.Database.Database
	if tenantId != "" {
		database = "tenant_" + tenantId
	}

	authServerURL := "http://localhost:" + strconv.Itoa(conf.Port)
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
	return
}

func NewServer(conf Config, tenantId string) *oauth {
	database := conf.Database.Database
	if tenantId != "" {
		database = "tenant_" + tenantId
	}
	authServerURL := "http://localhost:" + strconv.Itoa(conf.Port)

	oauthConfig := NewConf(conf, tenantId)

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	username := conf.Database.Username
	pwd := conf.Database.Password
	host := conf.Database.Host
	port := strconv.Itoa(conf.Database.Port)
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

	return &oauth{
		Config: oauthConfig,
		Srv:    srv,
		Store:  mysqlStore,
	}
}
