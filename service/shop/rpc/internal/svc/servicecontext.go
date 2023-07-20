package svc

import (
	"zerocmf/service/shop/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	tables := []string{
		"goods_category",
	}
	c.Database.Migrate(tables)
	return &ServiceContext{
		Config: c,
	}
}
