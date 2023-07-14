package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"
	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"
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

func (l *GetLogic) Get(in *user.UserRequest) (userReply *user.UserReply, err error) {
	c := l.svcCtx

	conf := c.Config.Database.NewConf(in.GetSiteId())
	db := database.NewGormDb(conf)

	id := in.GetUserId()

	query := "id = ?"
	var queryArgs interface{} = id

	userLogin := in.GetUserLogin()
	if userLogin != "" {
		query = "user_login = ?"
		queryArgs = userLogin
	}

	userModel := model.User{}
	tx := db.Where(query, queryArgs).First(&userModel)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	userReply = new(user.UserReply)
	copier.Copy(&userReply, &userModel)

	if userModel.Id == 0 {
		userReply.ErrorMsg = "该用户不存在或已被删除！"
	}

	userReply.UserPass = ""
	userReply.SiteId = in.GetSiteId()
	return
}
