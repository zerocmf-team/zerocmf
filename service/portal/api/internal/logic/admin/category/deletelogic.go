package category

import (
	"context"
	"net/http"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(header *http.Request, svcCtx *svc.ServiceContext) *DeleteLogic {
	ctx := header.Context()
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.CateOneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	PortalCategory := new(model.PortalCategory)
	PortalCategory.Id = id

	err := PortalCategory.Delete(db)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
