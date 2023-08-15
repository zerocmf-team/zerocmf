package product

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/productservice"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

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

	client := productservice.NewProductService(c.Client)

	rpcReq := productservice.ProductGetReq{
		SiteId:   siteId.(int64),
		Current:  &current,
		PageSize: &pageSize,
	}

	copier.Copy(&rpcReq, &req)

	products, getErr := client.ProductGet(ctx, &rpcReq)
	if getErr != nil {
		resp.Error("系统错误", getErr.Error())
		return
	}

	var json interface{} = products.GetData()
	if pageSize > 0 {
		paginate := data.Paginate{
			Current:  int(current),
			PageSize: int(pageSize),
			Data:     products.GetData(),
			Total:    products.GetTotal(),
		}
		if len(products.GetData()) == 0 {
			paginate.Data = make([]string, 0)
		}
		json = paginate
	}
	resp.Success("获取成功", json)

	return
}
