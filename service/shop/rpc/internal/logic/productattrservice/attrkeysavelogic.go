package productattrservicelogic

import (
	"context"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrKeySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttrKeySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttrKeySaveLogic {
	return &AttrKeySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttrKeySaveLogic) AttrKeySave(in *shop.ProductAttrKeyReq) (*shop.ProductAttrKeyResp, error) {
	// todo: add your logic here and delete this line

	return &shop.ProductAttrKeyResp{}, nil
}
