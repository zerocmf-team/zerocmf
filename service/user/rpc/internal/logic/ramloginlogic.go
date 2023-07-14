package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"

	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RamLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRamLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RamLoginLogic {
	return &RamLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RamLoginLogic) RamLogin(in *user.LoginReq) (userReply *user.UserReply, err error) {
	c := l.svcCtx

	conf := c.Config.Database.NewConf(in.GetSiteId())
	fmt.Println("conf", conf)
	db := database.NewGormDb(conf)

	userLogin := in.GetUserLogin()
	userPass := in.GetUserPass()

	query := "user_login = ?"
	queryArgs := userLogin

	userModel := model.User{}
	tx := db.Where(query, queryArgs).First(&userModel)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}

	userReply = new(user.UserReply)

	if userPass != userModel.UserPass {
		userReply.ErrorMsg = "账号密码不正确"
		return
	}

	copier.Copy(&userReply, &userModel)
	userReply.UserPass = ""
	userReply.SiteId = in.GetSiteId()
	return
}
