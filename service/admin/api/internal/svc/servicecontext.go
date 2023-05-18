package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/config"
	"zerocmf/service/tenant/rpc/tenantclient"
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
	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))

	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Db:             curDb,
		Data:           data,
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
	}
}
