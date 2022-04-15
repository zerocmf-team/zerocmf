package breadcrumb

import (
	"context"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type BreadcrumbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBreadcrumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BreadcrumbLogic {
	return &BreadcrumbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BreadcrumbLogic) Breadcrumb(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	id := req.Id

	if id == 0 {
		resp.Error("cid不能为空", nil)
		return
	}

	breadcrumbs, err := new(model.PortalCategory).GetPrevious(db, id)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功", breadcrumbs)
	return
}
