package category

import (
	"context"
	"net/http"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewEditLogic(header *http.Request, svcCtx *svc.ServiceContext) *EditLogic {
	ctx := header.Context()
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.CateSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = Save(c, req)
	return
}
