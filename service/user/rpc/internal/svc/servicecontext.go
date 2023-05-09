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

	curDb := database.NewDb(c.Database)
	// 设置为默认的db
	db := curDb.Db() // 初始化
	// 数据库迁移
	//model.Migrate(db)

	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
