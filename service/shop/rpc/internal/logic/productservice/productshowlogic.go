package productservicelogic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/model"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductShowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductShowLogic {
	return &ProductShowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductShowLogic) ProductShow(in *shop.ProductShowReq) (*shop.ProductResp, error) {

	ctx := l.ctx
	c := l.svcCtx

	conf := c.Config.Database.NewConf(in.GetSiteId())
	dsn := conf.Dsn()

	//mysql model调用
	conn := sqlx.NewMysql(dsn)

	productModel := model.NewProductModel(conn, c.Config.Cache)
	productId := in.GetProductId()
	one, err := productModel.FindOne(ctx, productId)
	if err != nil {
		return nil, err
	}

	resp := shop.ProductResp{}

	// 查询规格
	skuModel := model.NewProductSkuModel(conn, c.Config.Cache)
	var productSku []*model.ProductSku
	productSku, err = skuModel.Where("product_id = ?", productId).Find(ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	err = copier.Copy(&resp.ProductSku, &productSku)
	if err != nil {
		return nil, err
	}

	if one.Attributes != "" {
		var attributes []*shop.Attributes
		err = json.Unmarshal([]byte(one.Attributes), &attributes)
		if err != nil {
			return nil, err
		}
		resp.Attributes = attributes
	}

	if one.ProductThumbnail.Valid {
		var thumbnail []string
		err = json.Unmarshal([]byte(one.ProductThumbnail.String), &thumbnail)
		if err != nil {
			return nil, err
		}
		resp.ProductThumbnail = thumbnail
	}

	if one.CreatedAt > 0 {
		createdAt := one.CreatedAt
		resp.CreatedTime = time.Unix(createdAt, 0).Format(data.TimeLayout)
	}
	if one.UpdatedAt > 0 {
		updatedAt := one.UpdatedAt
		resp.UpdatedTime = time.Unix(updatedAt, 0).Format(data.TimeLayout)
	}

	err = copier.Copy(&resp, &one)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
