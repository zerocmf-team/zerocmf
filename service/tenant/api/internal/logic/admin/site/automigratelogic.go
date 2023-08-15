package site

import (
	"context"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/lowcode/rpc/lowcodeclient"
	"zerocmf/service/portal/rpc/portalclient"
	"zerocmf/service/shop/rpc/pb/shop"
	"zerocmf/service/tenant/model"
	"zerocmf/service/user/rpc/userclient"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoMigrateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoMigrateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoMigrateLogic {
	return &AutoMigrateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoMigrateLogic) AutoMigrate(req *types.SiteShowReq) (resp types.Response) {
	c := l.svcCtx
	userId, _ := c.Get("userId")
	db := c.Db
	context := l.ctx
	user := model.User{}
	tx := db.Where("id", userId).First(&user)
	if tx.Error != nil {
		resp.Error("操作失败！", tx.Error.Error())
		return
	}
	siteId := req.SiteId

	c.AdminRpc.AutoMigrate(context, &adminclient.SiteReq{
		SiteId: siteId,
	})

	c.UserRpc.AutoMigrate(context, &userclient.SiteReq{
		SiteId: siteId,
	})

	c.PortalRpc.AutoMigrate(context, &portalclient.SiteReq{
		SiteId: siteId,
	})

	c.LowcodeRpc.AutoMigrate(context, &lowcodeclient.SiteReq{
		SiteId: siteId,
	})

	c.ShopRpc.AutoMigrate(context, &shop.MigrateReq{
		SiteId: siteId,
	})

	resp.Success("操作成功！", nil)
	return
}
