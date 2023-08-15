package productservicelogic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/shop/model"

	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductGetLogic {
	return &ProductGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductGetLogic) ProductGet(in *shop.ProductGetReq) (*shop.ProductListResp, error) {

	ctx := l.ctx
	c := l.svcCtx

	conf := c.Config.Database.NewConf(in.GetSiteId())
	dsn := conf.Dsn()

	//mysql model调用
	conn := sqlx.NewMysql(dsn)

	productModel := model.NewProductModel(conn, c.Config.Cache)

	current := in.GetCurrent()
	pageSize := in.GetPageSize()

	var query []string
	var queryArgs []interface{}

	productName := in.GetProductName()

	if strings.TrimSpace(productName) != "" {
		query = append(query, "product_name like ?")
		queryArgs = append(queryArgs, "%"+productName+"%")
	}

	pid := in.GetProductCategory()
	if pid > 0 {
		query = append(query, "product_category = ?")
		queryArgs = append(queryArgs, pid)
	}

	status := in.Status
	if status != nil {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, in.GetStatus())
	}

	queryStr := strings.Join(query, " AND ")

	// 获取商品列表
	productList, err := productModel.Where(queryStr, queryArgs...).Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).OrderBy("product_id desc").Find(ctx)
	if err != nil {
		return nil, err
	}

	var total int64
	if pageSize > 0 {
		total, err = productModel.Count(ctx)
		if err != nil {
			return nil, err
		}
	}

	var categoryId = make([]string, 0)
	for _, v := range productList {
		cid := strconv.FormatInt(v.ProductCategory, 10)
		categoryId = append(categoryId, cid)
	}

	categoryId = util.RemoveDuplicates(categoryId)

	productCategoryModel := model.NewProductCategoryModel(conn, c.Config.Cache)

	var (
		productCategory []*model.ProductCategory
		categoryIdStr   string
		categoryQuery   string
	)

	if len(categoryId) > 0 {
		categoryIdStr = strings.Join(categoryId, ",")
		categoryQuery = fmt.Sprintf("product_category_id IN (%s)", categoryIdStr)
	}

	productCategory, err = productCategoryModel.Where(categoryQuery).Find(ctx)
	if err != nil {
		return nil, err
	}

	var productCategoryMap = make(map[int64]model.ProductCategory, len(productCategory))
	for _, v := range productCategory {
		productCategoryMap[v.ProductCategoryId] = *v
	}

	resp := shop.ProductListResp{}
	err = copier.Copy(&resp.Data, &productList)
	if err != nil {
		return nil, err
	}

	var _data = make([]*shop.ProductResp, len(productList))
	for k, v := range productList {

		item := shop.ProductResp{}
		if v.ProductCategory > 0 {
			item.ProductCategoryName = productCategoryMap[v.ProductCategory].Name
		}

		if v.ProductThumbnail.Valid {
			err = json.Unmarshal([]byte(v.ProductThumbnail.String), &item.ProductThumbnail)
			if err != nil {
				return nil, err
			}
		}

		if v.Attributes != "" {
			var attributes []*shop.Attributes
			err = json.Unmarshal([]byte(v.Attributes), &attributes)
			if err != nil {
				return nil, err
			}
			item.Attributes = attributes
		}

		err = copier.Copy(&item, &v)
		if err != nil {
			return nil, err
		}

		if v.PriceNegotiable.Valid {
			item.PriceNegotiable = v.PriceNegotiable.Int64
		}

		if v.CreatedAt > 0 {
			createdAt := v.CreatedAt
			item.CreatedTime = time.Unix(createdAt, 0).Format(data.TimeLayout)
		}
		if v.UpdatedAt > 0 {
			updatedAt := v.UpdatedAt
			item.UpdatedTime = time.Unix(updatedAt, 0).Format(data.TimeLayout)
		}
		_data[k] = &item

	}

	resp.Data = _data
	resp.Total = total
	return &resp, nil
}
