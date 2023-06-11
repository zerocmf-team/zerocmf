package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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
	NewDb   func(siteId ...string) database.GormDB
	Request *http.Request
	*Init.Data
	AuthMiddleware rest.Middleware
	UserRpc        user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 设置为默认的db

	// 数据库迁移
	//model.Migrate(db)
	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))
	return &ServiceContext{
		Config: c,
		NewDb: func(siteId ...string) (gormDB database.GormDB) {
			name := ""
			if len(siteId) > 0 {
				name = "site_" + siteId[0] + "_" + c.Name
			}
			c.Database.Database = name
			db := database.NewGormDb(c.Database)
			gormDB.Db = db
			gormDB.Database = c.Database
			return
		},
		Data:           data,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
	}
}
