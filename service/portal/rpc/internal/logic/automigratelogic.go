package logic

import (
	"context"
	"strconv"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/rpc/internal/svc"
	"zerocmf/service/portal/rpc/types/portal"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoMigrateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoMigrateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoMigrateLogic {
	return &AutoMigrateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AutoMigrateLogic) AutoMigrate(in *portal.SiteReq) (*portal.SiteReply, error) {
	c := l.svcCtx
	dbConf := c.Config.Database
	db := database.NewDb(dbConf)
	siteId := in.SiteId
	if siteId > 0 {
		// todo dsn 初始化
		siteStr := strconv.FormatInt(siteId, 10)
		dbORM := db.ManualDb(siteStr)
		model.Migrate(dbORM)
	}
	return &portal.SiteReply{}, nil
}
