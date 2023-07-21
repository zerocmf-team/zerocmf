package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/middleware"
	"zerocmf/service/shop/api/internal/config"
	"zerocmf/service/tenant/rpc/tenantclient"
)

type ServiceContext struct {
	Config         config.Config
	Client         zrpc.Client
	SiteMiddleware rest.Middleware
	AuthMiddleware rest.Middleware
	*Init.Data
}

func NewServiceContext(c config.Config) *ServiceContext {

	routes := []apisix.Route{
		{
			URI:       "/api/v1/shop/admin/*",
			Name:      "shop-api-admin",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				JWTAuth: &apisix.JWTAuth{
					Meta: apisix.Meta{
						Disable: false,
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
	client := zrpc.MustNewClient(c.ShopRpc)
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))
	return &ServiceContext{
		Config:         c,
		Client:         client,
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
		SiteMiddleware: middleware.NewSiteMiddleware(data).Handle,
	}
}
