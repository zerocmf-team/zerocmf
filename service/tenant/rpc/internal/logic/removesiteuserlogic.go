package logic

import (
	"context"
	"gorm.io/gorm"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveSiteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveSiteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveSiteUserLogic {
	return &RemoveSiteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveSiteUserLogic) RemoveSiteUser(in *tenant.RemoveSiteUserReq) (reply *tenant.RemoveSiteUserReply, err error) {

	reply = new(tenant.RemoveSiteUserReply)

	c := l.svcCtx
	db := c.Db

	user := model.User{}
	tx := db.Where("mobile = ? AND delete_at = 0", in.GetMobile()).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		err = tx.Error
		return
	}

	siteUser := model.SiteUser{}
	tx = db.Where("site_id = ? AND uid = ?", in.GetSiteId(), user.Uid).Delete(&siteUser)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}
