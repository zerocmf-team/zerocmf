package adminMenu

import (
	"context"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLogic {
	return &SyncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncLogic) Sync() (resp *types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	model.InitMenus(db)
	resp.Success("执行成功！", nil)
	return
}
