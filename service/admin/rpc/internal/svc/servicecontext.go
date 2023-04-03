package svc

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/model"
	"zerocmf/service/admin/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := database.NewDb(c.Database)
	// 数据库迁移
	curDb := database.Db()

	model.Migrate("", false)

	return &ServiceContext{
		Config: c,
		Db:     curDb,
	}
}
