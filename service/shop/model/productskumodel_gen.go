// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productSkuFieldNames          = builder.RawFieldNames(&ProductSku{})
	productSkuRows                = strings.Join(productSkuFieldNames, ",")
	productSkuRowsExpectAutoSet   = strings.Join(stringx.Remove(productSkuFieldNames, "`sku_id`", "`id`"), ",")
	productSkuRowsWithPlaceHolder = strings.Join(stringx.Remove(productSkuFieldNames, "`sku_id`", "`id`"), "=?,") + "=?"

	cacheProductSkuSkuIdPrefix = "cache:productSku:skuId:"
)

type (
	productSkuModel interface {
		Where(query string, args ...interface{}) *defaultProductSkuModel
		Limit(limit int) *defaultProductSkuModel
		Offset(offset int) *defaultProductSkuModel
		OrderBy(query string) *defaultProductSkuModel
		First(ctx context.Context) (*ProductSku, error)
		Find(ctx context.Context) ([]*ProductSku, error)
		Count(ctx context.Context) (int64, error)
		Insert(ctx context.Context, data *ProductSku) (sql.Result, error)
		FindOne(ctx context.Context, skuId int64) (*ProductSku, error)
		Update(ctx context.Context, data *ProductSku) error
		Delete(ctx context.Context, skuId int64) error
	}

	defaultProductSkuModel struct {
		sqlc.CachedConn
		table     string
		query     string
		queryArgs []interface{}
		limit     int
		offset    int
		orderBy   string
	}

	ProductSku struct {
		SkuId         int64           `db:"sku_id"`         // SKU的唯一标识符，主键，自增
		ProductId     int64           `db:"product_id"`     // 所属SPU的标识符，外键关联SPU表
		SkuCode       sql.NullString  `db:"sku_code"`       // 规格编码
		SkuBarcode    sql.NullString  `db:"sku_barcode"`    // 规格条码
		AttrsVal      string          `db:"attrs_val"`      // 属性值组合，例如"颜色:红色,存储容量:256G,网络类型:全网通"
		RetailPrice   float64         `db:"retail_price"`   // 零售价
		Stock         int64           `db:"stock"`          // 商品库存数量
		OriginalPrice sql.NullFloat64 `db:"original_price"` // 标准价
		CostPrice     sql.NullFloat64 `db:"cost_price"`     // 成本价
		Weight        sql.NullFloat64 `db:"weight"`         // 重量（g）
		Status        int64           `db:"status"`         // 状态（0 => 停用;1 => 启用）
		CreatedAt     int64           `db:"created_at"`     // 创建时间
		UpdatedAt     int64           `db:"updated_at"`     // 更新时间
		DeletedAt     int64           `db:"deleted_at"`     // 删除时间
	}
)

func newProductSkuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductSkuModel {
	return &defaultProductSkuModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`product_sku`",
	}
}

func (m *defaultProductSkuModel) withSession(session sqlx.Session) *defaultProductSkuModel {
	return &defaultProductSkuModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`product_sku`",
	}
}

func (m *defaultProductSkuModel) Where(query string, args ...interface{}) *defaultProductSkuModel {
	m.query = query
	m.queryArgs = args
	return m
}

func (m *defaultProductSkuModel) Limit(limit int) *defaultProductSkuModel {
	m.limit = limit
	return m
}

func (m *defaultProductSkuModel) Offset(offset int) *defaultProductSkuModel {
	m.offset = offset
	return m
}

func (m *defaultProductSkuModel) OrderBy(orderBy string) *defaultProductSkuModel {
	m.orderBy = orderBy
	return m
}
func (m *defaultProductSkuModel) Delete(ctx context.Context, skuId int64) error {
	productSkuSkuIdKey := fmt.Sprintf("%s%v", cacheProductSkuSkuIdPrefix, skuId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `sku_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, skuId)
	}, productSkuSkuIdKey)
	return err
}

func (m *defaultProductSkuModel) FindOne(ctx context.Context, skuId int64) (*ProductSku, error) {
	productSkuSkuIdKey := fmt.Sprintf("%s%v", cacheProductSkuSkuIdPrefix, skuId)
	var resp ProductSku
	err := m.QueryRowCtx(ctx, &resp, productSkuSkuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `sku_id` = ? AND deleted_at = 0 limit 1", productSkuRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, skuId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductSkuModel) Insert(ctx context.Context, data *ProductSku) (sql.Result, error) {
	productSkuSkuIdKey := fmt.Sprintf("%s%v", cacheProductSkuSkuIdPrefix, data.SkuId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, productSkuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.SkuCode, data.SkuBarcode, data.AttrsVal, data.RetailPrice, data.Stock, data.OriginalPrice, data.CostPrice, data.Weight, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	}, productSkuSkuIdKey)
	return ret, err
}

func (m *defaultProductSkuModel) Update(ctx context.Context, data *ProductSku) error {
	productSkuSkuIdKey := fmt.Sprintf("%s%v", cacheProductSkuSkuIdPrefix, data.SkuId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `sku_id` = ?", m.table, productSkuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductId, data.SkuCode, data.SkuBarcode, data.AttrsVal, data.RetailPrice, data.Stock, data.OriginalPrice, data.CostPrice, data.Weight, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.SkuId)
	}, productSkuSkuIdKey)
	return err
}

// 根据条件进行查询一条数据
func (m *defaultProductSkuModel) First(ctx context.Context) (*ProductSku, error) {
	query := m.query

	queryArgs := m.queryArgs
	orderBy := m.orderBy
	var resp ProductSku
	sql := fmt.Sprintf("select %s from %s", productSkuRows, m.table)

	if query != "" {
		sql += " where " + query
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
		if orderBy != "" {
			sql += fmt.Sprintf(" %s", orderBy)
		}
	}

	sql += " AND deleted_at = 0 limit 1"

	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 根据条件进行列表查询
func (m *defaultProductSkuModel) Find(ctx context.Context) ([]*ProductSku, error) {

	query := m.query
	queryArgs := m.queryArgs
	orderBy := m.orderBy

	var resp []*ProductSku
	sql := fmt.Sprintf("select %s from %s", productSkuRows, m.table)

	if query != "" {
		sql += " where " + query + " AND deleted_at = 0"
	} else {
		sql += " where deleted_at = 0"
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
		if orderBy != "" {
			sql += fmt.Sprintf(" %s", orderBy)
		}
	}

	limit := m.limit
	offset := m.offset

	// 查询条件
	if limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", limit)
	}

	if offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", offset)
	}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 统计字段
func (m *defaultProductSkuModel) Count(ctx context.Context) (int64, error) {
	query := m.query
	queryArgs := m.queryArgs
	sql := fmt.Sprintf("select count(`sku_id`) from %s", m.table)
	if query != "" {
		sql += " where " + query
	}
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
func (m *defaultProductSkuModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheProductSkuSkuIdPrefix, primary)
}

func (m *defaultProductSkuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `sku_id` = ? AND deleted_at = 0 limit 1", productSkuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductSkuModel) tableName() string {
	return m.table
}