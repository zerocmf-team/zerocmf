package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/redis"
	"zerocmf/service/lowcode/api/internal/config"
	"zerocmf/service/lowcode/api/internal/middleware"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	MongoDB func(dbName ...string) (db database.MongoDB, err error)
	Redis   func() redis.Redis
	Request *http.Request
	*Init.Data

	AuthMiddleware rest.Middleware
	SiteMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))

	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		MongoDB: func(dbName ...string) (db database.MongoDB, err error) {
			name := ""
			if len(dbName) > 0 {
				name = dbName[0]
			}
			db, err = database.NewMongoDB(c.MongoDB, name)
			return
		},

		Redis: func() redis.Redis {
			return redis.NewRedis(c.Redis)
		},
		Data:           data,
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
		SiteMiddleware: middleware.NewSiteMiddleware(data),
	}
}
