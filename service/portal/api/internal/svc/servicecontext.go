package svc

import (
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/db"
	"gincmf/service/portal/api/internal/config"
	"gincmf/service/portal/api/internal/middleware"
	"gincmf/service/portal/model"
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
	*data.Data
	AuthMiddleware rest.Middleware
	UserRpc        user.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := db.Conf()
	copier.Copy(&database, &c.Database)
	data.InitConfig(database)

	// 设置为默认的db
	db := database.Db() // 初始化

	// 数据库迁移
	model.Migrate("")
	data := new(data.Data)
	data.Context = data.InitContext()

	return &ServiceContext{
		Config:         c,
		Db:             db,
		Data:           data,
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
