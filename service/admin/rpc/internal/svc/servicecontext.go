package svc

import (
	"zerocmf/common/bootstrap/redis"
	"zerocmf/service/admin/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//Db     *gorm.DB
	Redis redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	//curDb := database.NewGormDb(c.Database)
	// 数据库迁移
	//model.Migrate(database.ManualDb("1161514444"))

	//model.Migrate(curDb)

	return &ServiceContext{
		Config: c,
		//Db:     curDb,
		Redis: redis.NewRedis(c.RedisConf),
	}
}
