package svc

import (
	"zerocmf/service/shop/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	tables := []string{
		"product",
		"product_resources",
		"product_attr_key",
		"product_attr_val",
		"product_sku",
		"product_sku_attr_relation",
		"product_category",
	}
	c.Database.Migrate(tables)
	return &ServiceContext{
		Config: c,
	}
}
