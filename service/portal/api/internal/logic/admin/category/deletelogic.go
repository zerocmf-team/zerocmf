package Categories

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

func (l *DeleteLogic) Delete(req *types.CateOneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	portalCategories := new(model.PortalCategories)
	portalCategories.Id = id

	err := portalCategories.Delete(db)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
