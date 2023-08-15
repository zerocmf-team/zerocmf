package category

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

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

func (l *ShowLogic) Show(req *types.CateOneReq) (resp types.Response) {
	id := req.Id
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	PortalCategory := new(model.PortalCategory)
	err := PortalCategory.Show(db, "id = ? and delete_at = ?", []interface{}{id, 0})
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", PortalCategory)
	return
}
