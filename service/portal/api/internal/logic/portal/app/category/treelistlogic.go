package category

import (
	"context"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"
	"github.com/zeromicro/go-zero/core/logx"
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
	db := c.Db
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	portalCategory := model.PortalCategory{
		ParentId: id,
	}

	trees, err := portalCategory.ListWithTree(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", trees)
	return
}
