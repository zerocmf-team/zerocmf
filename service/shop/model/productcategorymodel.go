package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductCategoryModel = (*customProductCategoryModel)(nil)

type (
	// ProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductCategoryModel.
	ProductCategoryModel interface {
		productCategoryModel
	}

	customProductCategoryModel struct {
		*defaultProductCategoryModel
	}
)

// NewProductCategoryModel returns a model for the database table.
func NewProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductCategoryModel {
	return &customProductCategoryModel{
		defaultProductCategoryModel: newProductCategoryModel(conn, c, opts...),
	}
}
