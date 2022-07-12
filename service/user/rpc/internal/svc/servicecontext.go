package svc

import (
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/user/model"
	"zerocmf/service/user/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db      *gorm.DB
	*data.Data
}

func NewServiceContext(c config.Config) *ServiceContext {
	curDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := curDb.Db() // 初始化
	model.Migrate("")

	return &ServiceContext{
		Config: c,
		Db: db,
	}
}
