package svc

import (
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/db"
	"gincmf/service/user/api/internal/config"
	"gincmf/service/user/model"
	"gincmf/service/user/rpc/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	*data.Data
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	database := db.Database()
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
	}
}
