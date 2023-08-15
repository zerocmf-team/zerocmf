package form

import (
	"context"
	"net/http"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(header *http.Request, svcCtx *svc.ServiceContext) *DeleteLogic {
	ctx := header.Context()
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.FormShowReq) (resp types.Response) {
	// todo: add your logic here and delete this line
	return
}
