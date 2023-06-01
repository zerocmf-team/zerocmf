package svc

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	*Init.Data
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := database.NewGormDb(c.Database)
	// 数据库迁移
	//model.Migrate(db)

	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
