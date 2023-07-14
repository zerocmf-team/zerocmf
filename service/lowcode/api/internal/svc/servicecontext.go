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

	routes := []apisix.Route{
		{
			URI:       "/api/v1/lowcode/admin/*",
			Name:      "lowcode-api-admin",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				JWTAuth: &apisix.JWTAuth{
					Meta: apisix.Meta{
						Disable: false,
					},
				},
				ProxyRewrite: &apisix.ProxyRewrite{
					RegexURI: []string{
						"^/api/v1/lowcode/admin/(.*)",
						"/api/v1/admin/$1",
					},
				},
			},
			Status: 1,
		},
		{
			URI:       "/api/v1/lowcode/app/*",
			Name:      "lowcode-api-app",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				ProxyRewrite: &apisix.ProxyRewrite{
					RegexURI: []string{
						"^/api/v1/lowcode/app/(.*)",
						"/api/v1/app/$1",
					},
				},
			},
			Status: 1,
		},
	}

	err := c.Apisix.Register(routes)
	if err != nil {
		panic(err)
	}

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
