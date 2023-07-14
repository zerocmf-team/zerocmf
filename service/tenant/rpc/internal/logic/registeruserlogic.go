package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strings"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/admin/rpc/types/admin"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 非主动创建站点

func (l *RegisterUserLogic) RegisterUser(in *tenant.RegisterReq) (reply *tenant.UserReply, err error) {

	reply = new(tenant.UserReply)

	c := l.svcCtx
	db := c.Db

	var uid int64

	query := make([]string, 0)
	queryArgs := make([]interface{}, 0)

	if in.GetMobile() != "" {
		query = append(query, "mobile = ?")
		queryArgs = append(queryArgs, in.GetMobile())
	}

	if in.GetUserLogin() != "" {
		query = append(query, "user_login = ?")
		queryArgs = append(queryArgs, in.GetUserLogin())
	}
	queryStr := strings.Join(query, " OR ")
	user := model.User{}
	tx := db.Where(queryStr, queryArgs...).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		err = errors.New("创建用户失败")
		return
	}

	if user.Id > 0 {
		/*uid = user.Uid
		user.DeleteAt = 0
		tx = db.Where("id = ?", user.Id).Updates(&user)*/
	} else {
		adminRpc := c.AdminRpc
		key := "tenant:user"
		var uidReply *adminclient.EncryptUidReply
		uidReply, err = adminRpc.EncryptUid(l.ctx, &admin.EncryptUidReq{Key: key, Salt: 0})
		if err != nil {
			return
		}
		uid = uidReply.Uid
		// 先创建统一站点用户
		user = model.User{
			Uid:       uid,
			UserPass:  in.GetUserPass(),
			UserLogin: in.UserLogin,
			Mobile:    in.GetMobile(),
		}
		tx = db.Create(&user)
	}

	if tx.Error != nil {
		err = errors.New("创建用户失败")
		return
	}

	// 创建关联站点信息
	siteUser := model.SiteUser{
		SiteId:  in.GetSiteId(),
		Uid:     uid,
		Oid:     in.GetOid(),
		IsOwner: 0,
	}
	tx = db.Create(&siteUser)
	if tx.Error != nil {
		return
	}
	reply.Id = user.Id
	reply.Uid = uid
	reply.Oid = in.GetOid()
	return
}
