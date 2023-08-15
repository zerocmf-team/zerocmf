package product

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/api/internal/types"
	"zerocmf/service/shop/rpc/client/productservice"
	"zerocmf/service/shop/rpc/pb/shop"

	"zerocmf/service/shop/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.ProductGetReq) (resp data.Rest) {

	ctx := l.ctx
	c := l.svcCtx

	siteId, _ := c.Get("siteId")

	current, pageSize, err := data.PaginateQueryInt32(l.header)
	if err != nil {
		resp.Error("系统错误", nil)
		return
	}

	productClient := productservice.NewProductService(c.Client)
	var (
		productListResp = new(productservice.ProductListResp)
	)

	productListResp, err = productClient.ProductGet(ctx, &productservice.ProductGetReq{
		SiteId:          siteId.(int64),
		Current:         &current,
		PageSize:        &pageSize,
		ProductCategory: req.ProductCategory,
	})
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	if productListResp.Data == nil {
		productListResp.Data = make([]*shop.ProductResp, 0)
	}

	var json interface{} = productListResp.Data

	if pageSize > 0 {
		json = data.Paginate{
			Current:  int(current),
			PageSize: int(pageSize),
			Data:     productListResp.Data,
			Total:    productListResp.GetTotal(),
		}
	}

	resp.Success("获取成功！", json)
	return
}
