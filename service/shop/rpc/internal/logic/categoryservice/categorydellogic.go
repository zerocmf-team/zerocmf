package categoryservicelogic

import (
	"context"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDelLogic {
	return &CategoryDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryDelLogic) CategoryDel(in *shop.CategoryDelReq) (*shop.CategoryResp, error) {
	// todo: add your logic here and delete this line

	return &shop.CategoryResp{}, nil
}
