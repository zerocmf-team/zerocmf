package logic

import (
	"context"
	"strconv"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/rpc/internal/svc"
	"zerocmf/service/admin/rpc/types/admin"

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

func (l *AutoMigrateLogic) AutoMigrate(in *admin.SiteReq) (*admin.SiteReply, error) {
	c := l.svcCtx
	dbConf := c.Config.Database
	dbORM := database.NewGormDb(dbConf)
	siteId := in.SiteId
	if siteId > 0 {
		// todo dsn 初始化
		siteStr := strconv.FormatInt(siteId, 10)
		dbORM = dbConf.ManualDb(siteStr)
		model.Migrate(dbORM)
	}
	return &admin.SiteReply{}, nil
}
