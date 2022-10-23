package category

import (
	"context"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.CateDelReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	portalCategory := new(model.PortalCategory)
	portalCategory.Id = id

	err := portalCategory.Delete(db)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
