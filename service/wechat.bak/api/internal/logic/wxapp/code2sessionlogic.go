package wxapp

import (
	"context"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Code2SessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCode2SessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Code2SessionLogic {
	return &Code2SessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Code2SessionLogic) Code2Session(req *types.Code2SessionReq) (resp types.Response) {
	// todo: add your logic here and delete this line

	return
}
