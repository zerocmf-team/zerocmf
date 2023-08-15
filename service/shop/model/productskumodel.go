package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSkuModel = (*customProductSkuModel)(nil)

type (
	// ProductSkuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductSkuModel.
	ProductSkuModel interface {
		productSkuModel
	}

	customProductSkuModel struct {
		*defaultProductSkuModel
	}
)

// NewProductSkuModel returns a model for the database table.
func NewProductSkuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductSkuModel {
	return &customProductSkuModel{
		defaultProductSkuModel: newProductSkuModel(conn, c, opts...),
	}
}

// NewProductSkuSessionModel returns a model for the Session database table.
func NewProductSkuSessionModel(conn sqlx.SqlConn, session sqlx.Session, c cache.CacheConf, opts ...cache.Option) ProductSkuModel {
	return &customProductSkuModel{
		defaultProductSkuModel: newProductSkuModel(conn, c, opts...).withSession(session),
	}
}
