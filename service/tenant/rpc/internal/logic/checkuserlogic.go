package logic

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserLogic {
	return &CheckUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserLogic) CheckUser(in *tenant.CheckUserReq) (reply *tenant.CheckUserReply, err error) {
	reply = new(tenant.CheckUserReply)

	c := l.svcCtx
	db := c.Db

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	queryOr := make([]string, 0)

	if in.GetMobile() != "" {
		queryOr = append(queryOr, "mobile = ?")
		queryArgs = append(queryArgs, in.GetMobile())
	}

	if in.GetUserLogin() != "" {
		queryOr = append(queryOr, "user_login = ?")
		queryArgs = append(queryArgs, in.GetUserLogin())
	}

	queryOrStr := strings.Join(queryOr, " AND ")

	query = append(query, queryOrStr)
	queryStr := strings.Join(queryOr, " OR ")

	// 根据手机号查询用户是否存在
	user := model.User{}
	tx := db.Where(queryStr, queryArgs...).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		err = tx.Error
		return
	}

	// 没查到用户
	if user.Id == 0 {
		return
	}

	siteUser := model.SiteUser{}
	tx = db.Where("site_id = ? AND uid = ?", in.GetSiteId(), user.Uid).First(&siteUser)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		err = tx.Error
		return
	}

	// 没绑定当前站点
	if siteUser.Id > 0 {
		reply.Uid = user.Uid
	}
	return
}
