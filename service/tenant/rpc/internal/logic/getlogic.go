package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *tenant.CurrentUserReq) (reply *tenant.UserReply, err error) {
	c := l.svcCtx
	db := c.Db
	uid := in.Uid
	siteId := in.SiteId
	if uid == "" {
		err = errors.New("参数错误")
		return
	}
	var user struct {
		SiteId int64 `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
		Oid    int64 `gorm:"type:bigint(20);comment:真实站点用户id;not null" json:"oid"`
	}
	prefix := c.Config.Database.Prefix
	tx := db.Select("s.site_id,su.oid,su.is_owner,su.list_order").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
		Joins("inner join "+prefix+"user u on u.uid = su.uid").
		Where("s.site_id = ? AND u.uid = ? AND s.delete_at = ?", siteId, uid, 0).Scan(&user)
	if tx.Error != nil {
		err = errors.New("用户不存在或已被删除！")
		return
	}
	reply = new(tenant.UserReply)
	copier.Copy(&reply, &user)
	return
}
