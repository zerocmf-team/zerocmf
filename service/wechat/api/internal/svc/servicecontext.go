package svc

import (
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/middleware"
	"zerocmf/common/bootstrap/redis"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/wechat/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  redis.Redis
	*Init.Data
	TenantRpc                      tenantclient.Tenant
	ComponentAccessTokenMiddleware rest.Middleware
	AuthMiddleware                 rest.Middleware
	SiteMiddleware                 rest.Middleware
	WxappMiddleware                rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	wxopen.NewOption(c.Wechat.WxOpen)

	routes := []apisix.Route{
		{
			URI:       "/api/v1/wechat/admin/*",
			Name:      "wechat-api-admin",
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
		{
			URI:       "/api/v1/wechat/*",
			Name:      "wechat-api",
			ServiceID: c.Apisix.Name,
			Status:    1,
		},
	}

	err := c.Apisix.Register(routes)
	if err != nil {
		panic(err)
	}

	data := new(Init.Data).Context()
	redis := redis.NewRedis(c.Redis)
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))

	return &ServiceContext{
		Config:                         c,
		Data:                           data,
		Redis:                          redis,
		TenantRpc:                      tenantRpc,
		ComponentAccessTokenMiddleware: middleware.ComponentAccessTokenMiddleware(data, redis),
		AuthMiddleware:                 apisix.AuthMiddleware(data, tenantRpc),
		SiteMiddleware:                 middleware.NewSiteMiddleware(data).Handle,
		WxappMiddleware:                middleware.NewWxappMiddleware(data, tenantRpc).Handle,
	}
}
