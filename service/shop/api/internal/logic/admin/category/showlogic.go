package category

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/categoryservice"

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

func (l *ShowLogic) Show(req *types.CategoryShowReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	categoryClient := categoryservice.NewCategoryService(c.Client)

	rpcReq := categoryservice.CategoryShowReq{}
	copier.Copy(&rpcReq, &req)

	category, err := categoryClient.CategoryShow(ctx, &rpcReq)
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}
	resp.Success("获取成功", category)
	return
}
