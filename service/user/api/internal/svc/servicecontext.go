package svc

import (
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/db"
	"gincmf/service/admin/rpc/admin"
	"gincmf/service/user/api/internal/config"
	"gincmf/service/user/api/internal/middleware"
	"gincmf/service/user/model"
	"gincmf/service/user/rpc/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	ResponseWriter http.ResponseWriter
	*data.Data
	UserRpc user.User
	AdminRpc admin.Admin
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := db.Conf()
	copier.Copy(&database, &c.Database)
	// 设置为默认的db
	db := database.Db() // 初始化
	// 数据库迁移
	model.Migrate("")
	data := new(data.Data)
	data.Context = data.InitContext()

	return &ServiceContext{
		Config:  c,
		Db:      db,
		Data:    data,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdminRpc:admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
