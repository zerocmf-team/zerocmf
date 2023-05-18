package svc

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/tenant/rpc/internal/config"
)

type ServiceContext struct {
	Db     *gorm.DB
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := database.NewDb(c.Database)
	// 数据库迁移
	curDb := database.Db()

	return &ServiceContext{
		Config: c,
		Db:     curDb,
	}
}
