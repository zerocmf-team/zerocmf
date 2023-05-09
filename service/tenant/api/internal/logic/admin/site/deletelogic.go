package site

import (
	"context"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.SiteShowReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	userId, _ := c.Get("userId")
	siteId := req.SiteId

	prefix := c.Config.Database.Prefix

	site := model.Site{}

	tx := db.Select("s.id,s.site_id,s.name,s.desc,s.status,s.create_at,su.oid,su.is_owner").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
		Joins("inner join "+prefix+"user u on u.uid = su.uid").
		Where("su.site_id = ? AND su.uid = ?", siteId, userId).
		First(&site)

	if util.IsDbErr(tx) != nil {
		resp.Error("操作失败！", tx.Error.Error())
		return
	}
	if tx.RowsAffected == 0 {
		resp.Error("操作失败！该站点不存在", tx.Error.Error())
		return
	}

	tx = db.Model(&site).Where("id", site.Id).Update("delete_at", time.Now().Unix())
	if tx.Error != nil {
		resp.Error("操作失败！", tx.Error.Error())
		return
	}
	resp.Success("删除成功！", &site)
	return
}
