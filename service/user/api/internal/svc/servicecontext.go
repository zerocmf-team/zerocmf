package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/rpc/adminclient"
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
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	curDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := curDb.Db() // 初始化
	// 数据库迁移
	//model.Migrate(db)

	data := new(Init.Data).Context()
	return &ServiceContext{
		Config:         c,
		Db:             db,
		Data:           data,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdminRpc:       adminclient.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		AuthMiddleware: apisix.AuthMiddleware(data),
	}
}
