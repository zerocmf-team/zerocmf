package logic

import (
	"context"
	"gincmf/service/user/model"
	"gincmf/service/user/rpc/internal/svc"
	"gincmf/service/user/rpc/types/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitLogic {
	return &InitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitLogic) Init(in *user.InitRequest) (*user.InitReply, error) {

	tenantId := in.GetTenantId()
	model.Migrate(tenantId)

	time.Sleep(time.Second * 5)

	return &user.InitReply{
		Code: 1,
	}, nil
}
