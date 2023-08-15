package productservicelogic

import (
	"context"
	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zerocmf/service/shop/model"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSaveLogic {
	return &ProductSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductSaveLogic) ProductSave(in *shop.ProductSaveReq) (*shop.ProductSaveResp, error) {

	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")

	//mysql model调用
	conn := sqlx.NewMysql(dsn)
	product := model.Product{}
	err := copier.Copy(&product, &in)
	if err != nil {
		return nil, err
	}

	productSku := in.GetProductSku()
	// 新增
	now := time.Now().Unix()
	if product.ProductId == 0 {

		product.CreatedAt = now
		product.UpdatedAt = now

		err = conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			// 新增商品
			var (
				insert    sql.Result
				productId int64
			)
			productModel := model.NewProductSessionModel(conn, session, c.Config.Cache)
			insert, err = productModel.Insert(ctx, &product)
			if err != nil {
				return err
			}

			productId, err = insert.LastInsertId()
			if err != nil {
				return err
			}
			product.ProductId = productId
			skuModel := model.NewProductSkuSessionModel(conn, session, c.Config.Cache)
			var skuItem *model.ProductSku
			for _, v := range productSku {
				skuItem, err = skuModel.Where("product_id = ? AND attrs_val = ?", []interface{}{productId, v.AttrsVal}...).First(ctx)
				if err != nil && err != model.ErrNotFound {
					return err
				}
				//	如果不存在则新增
				if skuItem == nil {
					skuItem = new(model.ProductSku)
					skuItem.CreatedAt = now
					skuItem.UpdatedAt = now
					err = copier.Copy(&skuItem, &v)
					if err != nil {
						return err
					}
					skuItem.ProductId = productId
					var (
						result sql.Result
					)
					result, err = skuModel.Insert(ctx, skuItem)
					if err != nil {
						return err
					}
					_, err = result.LastInsertId()
					if err != nil {
						return err
					}

				} else {
					err = copier.Copy(&skuItem, &v)
					if err != nil {
						return err
					}
					skuItem.ProductId = productId
					err = skuModel.Update(ctx, skuItem)
					if err != nil {
						return err
					}
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {

	}

	resp := shop.ProductSaveResp{
		ProductId:   product.ProductId,
		ProductName: product.ProductName,
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
