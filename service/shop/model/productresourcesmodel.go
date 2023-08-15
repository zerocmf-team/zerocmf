package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductResourcesModel = (*customProductResourcesModel)(nil)

type (
	// ProductResourcesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductResourcesModel.
	ProductResourcesModel interface {
		productResourcesModel
	}

	customProductResourcesModel struct {
		*defaultProductResourcesModel
	}
)

// NewProductResourcesModel returns a model for the database table.
func NewProductResourcesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductResourcesModel {
	return &customProductResourcesModel{
		defaultProductResourcesModel: newProductResourcesModel(conn, c, opts...),
	}
}
