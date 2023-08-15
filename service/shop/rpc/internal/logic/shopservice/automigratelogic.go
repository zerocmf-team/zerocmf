package shopservicelogic

import (
	"context"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

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

func (l *AutoMigrateLogic) AutoMigrate(in *shop.MigrateReq) (*shop.MigrateReply, error) {

	c := l.svcCtx

	config := c.Config

	tables := []string{
		"product",
		"product_resources",
		"product_attr_key",
		"product_attr_val",
		"product_sku",
		"product_sku_attr_relation",
		"product_category",
	}

	conf := config.Database.NewConf(in.SiteId)
	conf.Migrate(tables)

	return &shop.MigrateReply{}, nil
}
