package verify

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/wechat/api/internal/svc"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp data.Rest) {
	// todo: add your logic here and delete this line

	return
}
