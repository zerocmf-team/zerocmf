package miniapp

import (
	"context"
	"zerocmf/service/tenant/rpc/types/tenant"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.MiniappShowReq) (resp types.Response) {

	c := l.svcCtx
	siteId := req.SiteId

	reply, err := c.TenantRpc.GetSiteMp(l.ctx, &tenant.SiteMpReq{
		SiteId: siteId,
	})
	if err != nil {
		resp.Error("请求错误！", err.Error())
		return
	}

	resp.Success("获取成功！", reply.GetData())

	return
}
