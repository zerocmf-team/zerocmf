package oauth

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserLogic {
	return &CurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserLogic) CurrentUser() (resp types.Response) {
	c := l.svcCtx
	r := c.Request
	userId, _ := c.Get("userId")
	db := c.Db
	r.ParseForm()
	siteId := strings.Join(r.Form["siteId"], "")
	user := model.User{}

	tx := db.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		msg := "系统错误：" + tx.Error.Error()
		if tx.Error == gorm.ErrRecordNotFound {
			msg = "该管理员账号不存在"
		}
		resp.Error(msg, nil)
		return
	}
	user.Access = 1
	if siteId != "" {
		var siteUser struct {
			SiteId int64 `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
			Oid    int64 `gorm:"type:bigint(20);comment:真实站点用户id;not null" json:"oid"`
		}
		prefix := c.Config.Database.Prefix
		tx = db.Select("s.site_id,su.oid,su.is_owner,su.list_order").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
			Joins("inner join "+prefix+"user u on u.uid = su.uid").
			Where("s.site_id = ? AND u.uid = ? AND s.delete_at = ?", siteId, userId, 0).Scan(&siteUser)

		if tx.Error != nil {
			resp.Error("用户不存在或已被删除！", nil)
			return
		}

		if tx.RowsAffected == 0 {
			user.Access = 0
		}

	}

	resp.Success("获取成功！", user)
	return
}
