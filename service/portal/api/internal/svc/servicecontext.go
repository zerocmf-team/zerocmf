package svc

import (
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/portal/api/internal/config"
	"zerocmf/service/portal/api/internal/middleware"
	"zerocmf/service/portal/model"
	"zerocmf/service/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	*Init.Data
	AuthMiddleware rest.Middleware
	UserRpc        user.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := database.Conf()
	copier.Copy(&database, &c.Database)

	// 设置为默认的db
	db := database.Db() // 初始化

	// 数据库迁移
	model.Migrate("")

	return &ServiceContext{
		Config:         c,
		Db:             db,
		Data:           new(Init.Data).Context(),
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
