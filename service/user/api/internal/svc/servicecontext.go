package svc

import (
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/rpc/admin"
	"zerocmf/service/user/api/internal/config"
	"zerocmf/service/user/api/internal/middleware"
	"zerocmf/service/user/model"
	"zerocmf/service/user/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	ResponseWriter http.ResponseWriter
	*data.Data
	UserRpc user.User
	AdminRpc admin.Admin
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	curDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := curDb.Db() // 初始化
	// 数据库迁移
	model.Migrate("")

	return &ServiceContext{
		Config:  c,
		Db:      db,
		Data:    new(data.Data).InitContext(),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdminRpc:admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
