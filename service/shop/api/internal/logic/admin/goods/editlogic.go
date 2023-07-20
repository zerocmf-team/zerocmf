package goods

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

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

func (l *EditLogic) Edit(req *types.GoodsEditReq) (resp data.Rest) {
	// todo: add your logic here and delete this line

	return
}
