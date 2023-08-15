package categoryservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zerocmf/service/shop/model"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryShowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryShowLogic {
	return &CategoryShowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryShowLogic) CategoryShow(in *shop.CategoryShowReq) (*shop.CategoryResp, error) {
	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")
	//mysql model调用
	db := model.NewProductCategoryModel(sqlx.NewMysql(dsn), conf.Cache)
	id := in.GetId()

	one, err := db.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := shop.CategoryResp{}
	copier.Copy(&resp, &one)

	return &resp, nil
}
