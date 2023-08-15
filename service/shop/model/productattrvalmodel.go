package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductAttrValModel = (*customProductAttrValModel)(nil)

type (
	// ProductAttrValModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductAttrValModel.
	ProductAttrValModel interface {
		productAttrValModel
	}

	customProductAttrValModel struct {
		*defaultProductAttrValModel
	}
)

// NewProductAttrValModel returns a model for the database table.
func NewProductAttrValModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductAttrValModel {
	return &customProductAttrValModel{
		defaultProductAttrValModel: newProductAttrValModel(conn, c, opts...),
	}
}
