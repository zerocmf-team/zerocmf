package svc

import (
	"gincmf/common/bootstrap/data"
	"gincmf/common/bootstrap/db"
	"gincmf/service/admin/api/internal/config"
	"gincmf/service/admin/model"
	"github.com/jinzhu/copier"
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
	database := db.Conf()
	copier.Copy(&database, &c.Database)
	// 设置为默认的db

	// 数据库迁移
	curDb := db.Conf().ManualDb("")
	model.Migrate("",true)
	data := new(data.Data)
	data.Context = data.InitContext()

	return &ServiceContext{
		Config: c,
		Db:     curDb,
		Data:   data,
	}
}
