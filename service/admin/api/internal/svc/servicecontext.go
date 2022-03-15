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
	database := db.Database()
	copier.Copy(&database, &c.Database)
	// 设置为默认的db
	db := database.Db() // 初始化
	// 数据库迁移
	model.Migrate(db)
	data := new(data.Data)
	data.Context = data.InitContext()

	return &ServiceContext{
		Config: c,
		Db:     db,
		Data:   data,
	}
}
