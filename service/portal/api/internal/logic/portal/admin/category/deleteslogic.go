package category

import (
	"context"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesLogic {
	return &DeletesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesLogic) Deletes() (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	r := c.Request
	r.ParseForm()
	ids := r.Form["ids[]"]
	portalCategory := new(model.PortalCategory)
	err := portalCategory.BatchDelete(db, ids)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
