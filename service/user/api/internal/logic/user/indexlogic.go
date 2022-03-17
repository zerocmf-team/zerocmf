package user

import (
	"context"

	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) IndexLogic {
	return IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)

	h := map[string]string{
		"hello": "world",
	}

	resp.Success("user api 1.0", h)

	return
}
