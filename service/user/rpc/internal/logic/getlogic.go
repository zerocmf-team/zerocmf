package logic

import (
	"context"

	"gincmf/service/user/rpc/internal/svc"
	"gincmf/service/user/rpc/types"

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

func (l *GetLogic) Get(in *userclient.UserRequest) (userReply *userclient.UserReply,err error) {
	// todo: add your logic here and delete this line

	return &userclient.UserReply{}, nil
}
