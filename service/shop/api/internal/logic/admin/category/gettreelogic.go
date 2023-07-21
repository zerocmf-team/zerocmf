package category

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/api/internal/types"
	"zerocmf/service/shop/rpc/client/categoryservice"

	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/shop/api/internal/svc"
)

type GetTreeLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetTreeLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetTreeLogic {
	ctx := header.Context()
	return &GetTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetTreeLogic) GetTree(req *types.CategoryTreeDataReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	categoryClient := categoryservice.NewCategoryService(c.Client)

	rpcReq := categoryservice.CategoryTreeReq{}
	copier.Copy(&rpcReq, &req)

	treeResp, err := categoryClient.CategoryTree(ctx, &rpcReq)
	if err != nil {
		resp.Error("获取失败！", nil)
		return
	}
	resp.Success("获取成功！", treeResp.Data)
	return
}
