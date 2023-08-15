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
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`product_id`", "`id`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`product_id`", "`id`"), "=?,") + "=?"

	cacheProductProductIdPrefix = "cache:product:productId:"
)

type (
	productModel interface {
		Where(query string, args ...interface{}) *defaultProductModel
		Limit(limit int) *defaultProductModel
		Offset(offset int) *defaultProductModel
		OrderBy(query string) *defaultProductModel
		First(ctx context.Context) (*Product, error)
		Find(ctx context.Context) ([]*Product, error)
		Count(ctx context.Context) (int64, error)
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, productId int64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, productId int64) error
	}

	defaultProductModel struct {
		sqlc.CachedConn
		table     string
		query     string
		queryArgs []interface{}
		limit     int
		offset    int
		orderBy   string
	}

	Product struct {
		ProductId           int64           `db:"product_id"`            // 商品ID，主键，自增长整数
		ProductName         string          `db:"product_name"`          // 商品名称，不可为空
		UserId              int64           `db:"userId"`                // 创建人
		Attributes          string          `db:"attributes"`            // 规格属性选项
		ProductBarcode      string          `db:"product_barcode"`       // 商品条码，不可为空
		ProductCategory     int64           `db:"product_category"`      // 商品分类
		ProductThumbnail    sql.NullString  `db:"product_thumbnail"`     // 商品缩略图
		MainVideo           sql.NullString  `db:"main_video"`            // 主图视频，存储视频的URL或文件路径
		ExplanationVideo    sql.NullString  `db:"explanation_video"`     // 讲解视频，存储视频的URL或文件路径
		Price               float64         `db:"price"`                 // 商品售价
		PriceNegotiable     sql.NullInt64   `db:"price_negotiable"`      // 价格面议
		StockUnit           sql.NullString  `db:"stock_unit"`            // 库存单位
		Stock               sql.NullInt64   `db:"stock"`                 // 库存
		ShareDescription    sql.NullString  `db:"share_description"`     // 分享描述，用于在分享时显示的商品描述
		ProductSellingPoint sql.NullString  `db:"product_selling_point"` // 商品卖点，突出商品的特点
		OriginalPrice       sql.NullFloat64 `db:"original_price"`        // 划线价，商品的原价或划线价
		CostPrice           sql.NullFloat64 `db:"cost_price"`            // 成本价，商品的成本价
		HideRemainingStock  sql.NullInt64   `db:"hide_remaining_stock"`  // 商品详情不显示剩余件数
		DeliveryMethod      sql.NullInt64   `db:"delivery_method"`       // 配送方式
		ProductContent      sql.NullString  `db:"product_content"`       // 图文信息，存储商品的图文信息
		Status              int64           `db:"status"`                // 状态（0 => 停用;1 => 启用）
		CreatedAt           int64           `db:"created_at"`            // 创建时间
		UpdatedAt           int64           `db:"updated_at"`            // 更新时间
		DeletedAt           int64           `db:"deleted_at"`            // 删除时间
	}
)

func newProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`product`",
	}
}

func (m *defaultProductModel) withSession(session sqlx.Session) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`product`",
	}
}

func (m *defaultProductModel) Where(query string, args ...interface{}) *defaultProductModel {
	m.query = query
	m.queryArgs = args
	return m
}

func (m *defaultProductModel) Limit(limit int) *defaultProductModel {
	m.limit = limit
	return m
}

func (m *defaultProductModel) Offset(offset int) *defaultProductModel {
	m.offset = offset
	return m
}

func (m *defaultProductModel) OrderBy(orderBy string) *defaultProductModel {
	m.orderBy = orderBy
	return m
}
func (m *defaultProductModel) Delete(ctx context.Context, productId int64) error {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, productId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `product_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, productId)
	}, productProductIdKey)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, productId int64) (*Product, error) {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, productId)
	var resp Product
	err := m.QueryRowCtx(ctx, &resp, productProductIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `product_id` = ? AND deleted_at = 0 limit 1", productRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, productId)
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

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, data.ProductId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductName, data.UserId, data.Attributes, data.ProductBarcode, data.ProductCategory, data.ProductThumbnail, data.MainVideo, data.ExplanationVideo, data.Price, data.PriceNegotiable, data.StockUnit, data.Stock, data.ShareDescription, data.ProductSellingPoint, data.OriginalPrice, data.CostPrice, data.HideRemainingStock, data.DeliveryMethod, data.ProductContent, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	}, productProductIdKey)
	return ret, err
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, data.ProductId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `product_id` = ?", m.table, productRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductName, data.UserId, data.Attributes, data.ProductBarcode, data.ProductCategory, data.ProductThumbnail, data.MainVideo, data.ExplanationVideo, data.Price, data.PriceNegotiable, data.StockUnit, data.Stock, data.ShareDescription, data.ProductSellingPoint, data.OriginalPrice, data.CostPrice, data.HideRemainingStock, data.DeliveryMethod, data.ProductContent, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.ProductId)
	}, productProductIdKey)
	return err
}

// 根据条件进行查询一条数据
func (m *defaultProductModel) First(ctx context.Context) (*Product, error) {
	query := m.query

	queryArgs := m.queryArgs
	orderBy := m.orderBy
	var resp Product
	sql := fmt.Sprintf("select %s from %s", productRows, m.table)

	if query != "" {
		sql += " where " + query
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
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
func (m *defaultProductModel) Find(ctx context.Context) ([]*Product, error) {

	query := m.query
	queryArgs := m.queryArgs
	orderBy := m.orderBy

	var resp []*Product
	sql := fmt.Sprintf("select %s from %s", productRows, m.table)

	if query != "" {
		sql += " where " + query + " AND deleted_at = 0"
	} else {
		sql += " where deleted_at = 0"
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
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
func (m *defaultProductModel) Count(ctx context.Context) (int64, error) {
	query := m.query
	queryArgs := m.queryArgs
	sql := fmt.Sprintf("select count(`product_id`) from %s", m.table)
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
func (m *defaultProductModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheProductProductIdPrefix, primary)
}

func (m *defaultProductModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `product_id` = ? AND deleted_at = 0 limit 1", productRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductModel) tableName() string {
	return m.table
}