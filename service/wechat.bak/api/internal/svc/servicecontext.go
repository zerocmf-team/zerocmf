package svc

import (
	// goRedis "github.com/go-redis/redis"
	// goSessions "github.com/gorilla/sessions"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
	"zerocmf/common/bootstrap/sessions"
	weData "zerocmf/service/wechat/api/data"
	"zerocmf/service/wechat/api/internal/config"
	// "zerocmf/service/wechat/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	Db             *gorm.DB
	Redis          *goRedis.Client
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Store          *goSessions.CookieStore
	WechatMpToken  rest.Middleware
	weData.MpInfo
	*Init.Data
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := database.NewGormDb(c.Database)
	// autoMigrate
	// model.Migrate("")
	redis := redis.NewRedis(c.Redis)
	store := sessions.NewStore()

	data := new(Init.Data).Context()
	// mpTokenMiddleware := middleware.NewWechatMpTokenMiddleware(data, redis)

	return &ServiceContext{
		Config:        c,
		Db:            db,
		Redis:         redis,
		Store:         store,
		// WechatMpToken: mpTokenMiddleware.Handle,
		Data:          data,
	}

}
