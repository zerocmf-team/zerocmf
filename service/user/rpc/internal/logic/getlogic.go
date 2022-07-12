package logic

import (
	"context"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"
	"github.com/jinzhu/copier"

	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"

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

func (l *GetLogic) Get(in *user.UserRequest) (userReply *user.UserReply,err error) {
	// todo: add your logic here and delete this line

	c := l.svcCtx
	db := c.Db

	id := in.GetUserId()

	userModel := model.User{}
	tx := db.Where("id = ?",id).First(&userModel)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}

	userReply = new(user.UserReply)
	copier.Copy(&userReply,&userModel)

	return
}
