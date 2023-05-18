package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/portal/api/internal/config"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/user/rpc/types/user"
	"zerocmf/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	*Init.Data
	AuthMiddleware rest.Middleware
	UserRpc        user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	newDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := newDb.Db() // 初始化
	// 数据库迁移
	//model.Migrate(db)

	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))
	return &ServiceContext{
		Config:         c,
		Db:             db,
		Data:           data,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
	}
}
