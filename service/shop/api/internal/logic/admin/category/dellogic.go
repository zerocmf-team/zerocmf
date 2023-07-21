package category

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/categoryservice"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDelLogic(header *http.Request, svcCtx *svc.ServiceContext) *DelLogic {
	ctx := header.Context()
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.CategoryDelReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	categoryClient := categoryservice.NewCategoryService(c.Client)
	rpcReq := categoryservice.CategoryDelReq{
		Id: req.Id,
	}
	category, err := categoryClient.CategoryDel(ctx, &rpcReq)
	if err != nil {
		resp.Error("删除失败！", err.Error())
		return
	}
	resp.Success("删除成功！", category)
	return
}
