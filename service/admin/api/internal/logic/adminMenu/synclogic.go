package adminMenu

import (
	"context"
	"gincmf/service/admin/model"

	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"

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

func (l *SyncLogic) Sync() (resp *types.Response, err error) {
	c := l.svcCtx
	db := c.Db
	model.InitMenus(db)
	resp.Success("执行成功！", nil)
	return
}
