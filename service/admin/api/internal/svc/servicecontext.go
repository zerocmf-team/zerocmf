package svc

import (
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/config"
	"zerocmf/service/admin/model"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	*data.Data
}

func NewServiceContext(c config.Config) *ServiceContext {
	database := database.NewDb(c.Database)
	// 数据库迁移
	curDb := database.Db()
	model.Migrate("",true)

	return &ServiceContext{
		Config: c,
		Db:     curDb,
		Data:   new(data.Data).InitContext(),
	}
}
