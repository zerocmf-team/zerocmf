package product

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/categoryservice"
	"zerocmf/service/shop/rpc/client/productservice"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/shop/api/internal/svc"
)

type TreeLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

type listResp struct {
	Category struct {
		ProductCategoryId int64  `json:"productCategoryId"`
		CategoryName      string `json:"categoryName"`
	} `json:"category"`
	Data []*shop.ProductResp `json:"data"`
}

func NewTreeLogic(header *http.Request, svcCtx *svc.ServiceContext) *TreeLogic {
	ctx := header.Context()
	return &TreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *TreeLogic) Tree() (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx

	siteId, _ := c.Get("siteId")
	categoryClient := categoryservice.NewCategoryService(c.Client)

	var pageSize int32 = 0
	rpcReq := categoryservice.CategoryGetReq{
		SiteId:   siteId.(int64),
		PageSize: &pageSize,
	}

	category, err := categoryClient.CategoryGet(ctx, &rpcReq)
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	productClient := productservice.NewProductService(c.Client)
	var productListResp *productservice.ProductListResp
	productListResp, err = productClient.ProductGet(ctx, &productservice.ProductGetReq{
		SiteId:   siteId.(int64),
		PageSize: &pageSize,
	})
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	productData := productListResp.GetData()
	var productListMap = make(map[int64][]*shop.ProductResp, len(productData))
	for _, v := range productData {
		categoryId := v.ProductCategory
		productListMap[categoryId] = append(productListMap[categoryId], v)
	}

	categoryData := category.GetData()

	list := make([]listResp, len(categoryData))

	for k, v := range categoryData {
		list[k].Category.ProductCategoryId = v.GetProductCategoryId()
		list[k].Category.CategoryName = v.GetName()
		categoryId := v.ProductCategoryId
		list[k].Data = productListMap[categoryId]
	}

	resp.Success("获取成功！", list)

	return
}
