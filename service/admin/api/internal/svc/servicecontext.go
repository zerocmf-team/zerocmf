package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/config"
	"zerocmf/service/admin/api/internal/middleware"
	"zerocmf/service/admin/model"
	"zerocmf/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	Db      *gorm.DB
	Request *http.Request
	*Init.Data
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := database.NewDb(c.Database)
	// 数据库迁移
	curDb := database.Db()
	model.Migrate("", true)

	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Db:             curDb,
		Data:           new(Init.Data).Context(),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
