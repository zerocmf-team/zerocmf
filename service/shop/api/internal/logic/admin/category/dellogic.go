package category

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDelLogic(header *http.Request, svcCtx *svc.ServiceContext) *DelLogic {
	ctx := header.Context()
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.CategoryDelReq) (resp data.Rest) {
	// todo: add your logic here and delete this line

	return
}
