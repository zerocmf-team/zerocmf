package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/portal/rpc/portalclient"
	"zerocmf/service/tenant/api/internal/config"
	"zerocmf/service/tenant/model"
	"zerocmf/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Redis   func() redis.Redis
	Request *http.Request
	*Init.Data
	AdminRpc       adminclient.Admin
	UserRpc        userclient.User
	PortalRpc      portalclient.Portal
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 设置为默认的db
	db := database.NewGormDb(c.Database)
	// 数据库迁移
	model.Migrate(db)

	data := new(Init.Data).Context()
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis: func() redis.Redis {
			return redis.NewRedis(c.Redis)
		},
		Data:           data,
		AdminRpc:       adminclient.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		PortalRpc:      portalclient.NewPortal(zrpc.MustNewClient(c.PortalRpc)),
		AuthMiddleware: apisix.AuthMiddleware(data, nil),
	}
}
