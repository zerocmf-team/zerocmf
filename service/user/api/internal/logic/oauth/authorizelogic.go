package oauth

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
)

type AuthorizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizeLogic {
	return &AuthorizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizeLogic) Authorize() (resp types.Response) {
	return
}
