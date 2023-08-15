package category

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeListLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewTreeListLogic(header *http.Request, svcCtx *svc.ServiceContext) *TreeListLogic {
	ctx := header.Context()
	return &TreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *TreeListLogic) TreeList(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	PortalCategory := model.PortalCategory{
		ParentId: id,
	}

	trees, err := PortalCategory.ListWithTree(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", trees)
	return
}
