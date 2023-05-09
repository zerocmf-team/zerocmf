package svc

import (
	goRedis "github.com/go-redis/redis"
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
	Redis   *goRedis.Client
	Request *http.Request
	*Init.Data
	AdminRpc       adminclient.Admin
	UserRpc        userclient.User
	PortalRpc      portalclient.Portal
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	curDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := curDb.Db() // 初始化
	// 数据库迁移
	model.Migrate("")
	client := redis.NewRedis(c.Redis)
	data := new(Init.Data).Context()
	return &ServiceContext{
		Config:         c,
		Db:             db,
		Redis:          client,
		Data:           data,
		AdminRpc:       adminclient.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		PortalRpc:      portalclient.NewPortal(zrpc.MustNewClient(c.PortalRpc)),
		AuthMiddleware: apisix.AuthMiddleware(data),
	}
}
