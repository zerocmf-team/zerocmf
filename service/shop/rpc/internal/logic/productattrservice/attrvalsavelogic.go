package productattrservicelogic

import (
	"context"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrValSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttrValSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttrValSaveLogic {
	return &AttrValSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttrValSaveLogic) AttrValSave(in *shop.ProductAttrValReq) (*shop.ProductAttrValResp, error) {
	// todo: add your logic here and delete this line

	return &shop.ProductAttrValResp{}, nil
}
