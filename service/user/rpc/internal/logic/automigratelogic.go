package logic

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/user/model"

	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"

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

func (l *AutoMigrateLogic) AutoMigrate(in *user.SiteReq) (*user.SiteReply, error) {
	c := l.svcCtx
	dbConf := c.Config.Database
	var dbORM *gorm.DB
	siteId := in.SiteId
	if siteId > 0 {
		// todo dsn 初始化
		siteStr := strconv.FormatInt(siteId, 10)
		dbORM = database.CreateGormDb(dbConf, siteStr)
	} else {
		dbORM = database.CreateGormDb(dbConf)
	}
	model.Migrate(dbORM)
	return &user.SiteReply{}, nil
}
