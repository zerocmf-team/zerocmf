package article

import (
	"context"
	"time"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

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
	portalPost := new(model.PortalPost)
	if err := db.Model(&portalPost).Where("id IN (?)", ids).Updates(map[string]interface{}{"delete_at": time.Now().Unix()}).Error; err != nil {
		resp.Error("删除失败！", nil)
		return
	}
	resp.Success("删除成功！", ids)
	return
}
