package category

import (
	"context"
	"net/http"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDeletesLogic(header *http.Request, svcCtx *svc.ServiceContext) *DeletesLogic {
	ctx := header.Context()
	return &DeletesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DeletesLogic) Deletes() (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := l.header
	r.ParseForm()
	ids := r.Form["ids[]"]
	PortalCategory := new(model.PortalCategory)
	err := PortalCategory.BatchDelete(db, ids)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
