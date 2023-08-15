package app

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewShowLogic(header *http.Request, svcCtx *svc.ServiceContext) *ShowLogic {
	ctx := header.Context()
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.AppShowReq) (resp types.Response) {
	return
}
