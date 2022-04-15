package svc

import (
	"gincmf/common/bootstrap/db"
	"gincmf/service/admin/model"
	"gincmf/service/admin/rpc/internal/config"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := db.Conf()
	copier.Copy(&database, &c.Database)
	// 设置为默认的db
	db := database.Db() // 初始化
	model.Migrate("",false)

	return &ServiceContext{
		Config: c,
		Db: db,
	}
}
