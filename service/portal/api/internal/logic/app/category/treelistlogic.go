package Categories

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
)

type TreeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeListLogic {
	return &TreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeListLogic) TreeList(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	portalCategories := model.PortalCategories{
		ParentId: id,
	}

	trees, err := portalCategories.ListWithTree(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", trees)
	return
}
