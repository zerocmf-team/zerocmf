package oauth

import (
	"context"
	"strconv"
	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"
	"zerocmf/service/tenant/model"
	userModel "zerocmf/service/user/model"
	"zerocmf/service/user/rpc/userclient"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

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
	tenantUserId, _ := c.Get("tenantUserId")
	userId, _ := c.Get("userId")
	loginType, _ := c.Get("loginType")
	siteId, exist := c.Get("siteId")

	userRpc := c.UserRpc

	var userModel struct {
		userModel.User
		SiteId    int64  `json:"site_id"`
		LoginType string `json:"login_type"`
	}

	if loginType == "ram" {
		userReply, err := userRpc.Get(l.ctx, &userclient.UserRequest{
			UserId: userId.(string),
			SiteId: siteId.(int64),
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		if userReply.ErrorMsg != "" {
			resp.Error(userReply.ErrorMsg, nil)
			return
		}

		copier.Copy(&userModel, &userReply)
		userModel.SiteId = userReply.SiteId
		userModel.LoginType = "ram"

		resp.Success("获取成功！", userModel)
		return
	}

	db := c.Db

	user := model.User{}

	tx := db.Where("uid = ?", userId).First(&user)
	if tx.Error != nil {
		msg := "系统错误：" + tx.Error.Error()
		if tx.Error == gorm.ErrRecordNotFound {
			msg = "该管理员账号不存在"
		}
		resp.Error(msg, nil)
		return
	}

	if exist {
		var siteUser struct {
			SiteId int64 `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
			Oid    int64 `gorm:"type:bigint(20);comment:真实站点用户id;not null" json:"oid"`
		}

		prefix := c.Config.Database.Prefix
		tx = db.Select("s.site_id,su.oid,su.is_owner,su.list_order").Table(prefix+"site s").Joins("left join "+prefix+"site_user su on s.site_id = su.site_id").
			Joins("inner join "+prefix+"user u on u.uid = su.uid").
			Where("s.site_id = ? AND u.uid = ? AND s.delete_at = ?", siteId, tenantUserId, 0).Scan(&siteUser)

		if tx.Error != nil {
			resp.Error("该用户不存在或已被删除！", nil)
			return
		}

		if siteUser.Oid == 0 {
			resp.Error("该站点暂无访问权限！", nil)
			return
		}

		userReply, err := userRpc.Get(l.ctx, &userclient.UserRequest{
			UserId: strconv.FormatInt(siteUser.Oid, 10),
			SiteId: siteId.(int64),
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		if userReply.ErrorMsg != "" {
			resp.Error(userReply.ErrorMsg, nil)
			return
		}

		copier.Copy(&userModel, &userReply)
		userModel.SiteId = userReply.SiteId
		resp.Success("获取成功！", userModel)
		return
	}

	resp.Success("获取成功！", user)
	return
}
