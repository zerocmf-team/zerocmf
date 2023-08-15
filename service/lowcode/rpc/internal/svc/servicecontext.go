package svc

import (
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/lowcode/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	MongoDB func(dbName ...int64) (db database.MongoDB, err error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MongoDB: func(dbName ...int64) (db database.MongoDB, err error) {
			var siteId int64 = 0
			if len(dbName) > 0 {
				siteId = dbName[0]
			}
			db, err = database.NewMongoDB(c.MongoDB, siteId)
			return
		},
	}
}
