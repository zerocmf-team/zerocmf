package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductAttrKeyModel = (*customProductAttrKeyModel)(nil)

type (
	// ProductAttrKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductAttrKeyModel.
	ProductAttrKeyModel interface {
		productAttrKeyModel
	}

	customProductAttrKeyModel struct {
		*defaultProductAttrKeyModel
	}
)

// NewProductAttrKeyModel returns a model for the database table.
func NewProductAttrKeyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductAttrKeyModel {
	return &customProductAttrKeyModel{
		defaultProductAttrKeyModel: newProductAttrKeyModel(conn, c, opts...),
	}
}
