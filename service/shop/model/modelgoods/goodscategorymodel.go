package modelgoods

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsCategoryModel = (*customGoodsCategoryModel)(nil)

type (
	// GoodsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsCategoryModel.
	GoodsCategoryModel interface {
		goodsCategoryModel
	}

	customGoodsCategoryModel struct {
		*defaultGoodsCategoryModel
	}
)

// NewGoodsCategoryModel returns a model for the database table.
func NewGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GoodsCategoryModel {
	return &customGoodsCategoryModel{
		defaultGoodsCategoryModel: newGoodsCategoryModel(conn, c, opts...),
	}
}
