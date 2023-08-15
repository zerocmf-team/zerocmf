package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSkuAttrRelationModel = (*customProductSkuAttrRelationModel)(nil)

type (
	// ProductSkuAttrRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductSkuAttrRelationModel.
	ProductSkuAttrRelationModel interface {
		productSkuAttrRelationModel
	}

	customProductSkuAttrRelationModel struct {
		*defaultProductSkuAttrRelationModel
	}
)

// NewProductSkuAttrRelationModel returns a model for the database table.
func NewProductSkuAttrRelationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductSkuAttrRelationModel {
	return &customProductSkuAttrRelationModel{
		defaultProductSkuAttrRelationModel: newProductSkuAttrRelationModel(conn, c, opts...),
	}
}
