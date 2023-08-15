package product

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDelLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewBatchDelLogic(header *http.Request, svcCtx *svc.ServiceContext) *BatchDelLogic {
	ctx := header.Context()
	return &BatchDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *BatchDelLogic) BatchDel(req *types.ProductBatchDelReq) (resp data.Rest) {
	// todo: add your logic here and delete this line

	return
}
