package svc

import (
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/db"
	"zerocmf/service/wechat/api/internal/config"
	"zerocmf/service/wechat/model"
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
	db := database.Db() // 初始化
	// autoMigrate
	model.Migrate("")

	return &ServiceContext{
		Config: c,
		Db:     db,
		Data:   new(data.Data).InitContext(),
	}

}
