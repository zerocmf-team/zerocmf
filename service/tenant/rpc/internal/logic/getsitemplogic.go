package logic

import (
	"context"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteMpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSiteMpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteMpLogic {
	return &GetSiteMpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSiteMpLogic) GetSiteMp(in *tenant.SiteMpReq) (*tenant.SiteMpReply, error) {

	c := l.svcCtx
	db := c.Db
	var mpAuth []model.SiteMpAuth
	tx := db.Where("site_id = ?", in.SiteId).Find(&mpAuth)
	if util.IsDbErr(tx) != nil {
		return &tenant.SiteMpReply{}, tx.Error
	}

	reply := tenant.SiteMpReply{}
	for _, v := range mpAuth {
		item := new(tenant.SiteMpData)
		item.SiteId = v.SiteId
		item.AuthAppId = v.AuthAppId
		item.Type = v.Type
		reply.Data = append(reply.Data, item)
	}
	return &reply, nil
}
