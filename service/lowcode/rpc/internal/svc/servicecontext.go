package svc

import (
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/lowcode/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	MongoDB func(dbName ...string) (db database.MongoDB, err error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MongoDB: func(dbName ...string) (db database.MongoDB, err error) {
			name := ""
			if len(dbName) > 0 {
				name = dbName[0]
			}
			db, err = database.NewMongoDB(c.MongoDB, name)
			return
		},
	}
}
