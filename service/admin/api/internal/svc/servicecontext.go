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

	routes := []apisix.Route{
		{
			URI:       "/api/v1/admin/*",
			Name:      "admin-api",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				JWTAuth: &apisix.JWTAuth{
					Meta: apisix.Meta{
						Disable: false,
					},
				},
				ProxyRewrite: &apisix.ProxyRewrite{
					RegexURI: []string{
						"^/api/v1/admin(.*)",
						"/api/v1$1",
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

	db := database.NewGormDb(c.Database)
	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))

	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Db:             db,
		Data:           data,
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
	}
}
