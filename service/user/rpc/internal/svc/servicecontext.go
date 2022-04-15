package svc

import (
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/db"
	"gincmf/service/user/model"
	"gincmf/service/user/rpc/internal/config"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db      *gorm.DB
	*data.Data
}

func NewServiceContext(c config.Config) *ServiceContext {
	database := db.Conf()
	copier.Copy(&database, &c.Database)
	// 设置为默认的db
	db := database.Db() // 初始化
	model.Migrate("")

	return &ServiceContext{
		Config: c,
		Db: db,
	}
}
