package product

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/productservice"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewShowLogic(header *http.Request, svcCtx *svc.ServiceContext) *ShowLogic {
	ctx := header.Context()
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ProductShowReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	siteId, _ := c.Get("siteId")

	client := productservice.NewProductService(c.Client)

	rpcReq := productservice.ProductShowReq{
		SiteId:    siteId.(int64),
		ProductId: req.ProductId,
	}

	show, err := client.ProductShow(ctx, &rpcReq)
	if err != nil {
		resp.Error("查询失败！", err.Error())
		return
	}

	resp.Success("获取成功！", show)
	return
}
