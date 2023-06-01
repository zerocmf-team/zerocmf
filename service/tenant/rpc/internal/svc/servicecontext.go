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

	db := database.NewGormDb(c.Database)
	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
