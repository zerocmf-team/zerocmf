package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/user/api/internal/config"
	"zerocmf/service/user/rpc/types/user"
	"zerocmf/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	Db             *gorm.DB
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	*Init.Data
	UserRpc        user.UserClient
	AdminRpc       adminclient.Admin
	TenantRpc      tenantclient.Tenant
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	routes := []apisix.Route{
		{
			URI:       "/api/v1/user/admin/*",
			Name:      "user-api-admin",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				JWTAuth: &apisix.JWTAuth{
					Meta: apisix.Meta{
						Disable: false,
					},
				},
				ProxyRewrite: &apisix.ProxyRewrite{
					RegexURI: []string{
						"^/api/v1/user/admin/(.*)",
						"/api/v1/admin/$1",
					},
				},
			},
			Status: 1,
		},
		{
			URI:       "/api/v1/user/app/*",
			Name:      "user-api-app",
			ServiceID: c.Apisix.Name,
			Plugins: apisix.RoutePlugins{
				ProxyRewrite: &apisix.ProxyRewrite{
					RegexURI: []string{
						"^/api/v1/user/app/(.*)",
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

	//db := database.NewGormDb(c.Database)
	// 设置为默认的db
	// 数据库迁
	//model.Migrate(db)

	data := new(Init.Data).Context()
	tenantRpc := tenantclient.NewTenant(zrpc.MustNewClient(c.TenantRpc))
	return &ServiceContext{
		Config: c,
		//Db:             db,
		Data:           data,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdminRpc:       adminclient.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		TenantRpc:      tenantRpc,
		AuthMiddleware: apisix.AuthMiddleware(data, tenantRpc),
	}
}
