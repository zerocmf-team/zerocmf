package categoryservicelogic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zerocmf/service/shop/model/modelgoods"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryGetLogic {
	return &CategoryGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryGetLogic) CategoryGet(in *shop.CategoryGetReq) (*shop.CategoryListResp, error) {
	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")
	//mysql model调用
	db := modelgoods.NewGoodsCategoryModel(sqlx.NewMysql(dsn), conf.Cache)

	current := in.GetCurrent()
	pageSize := in.GetPageSize()

	list, err := db.Limit(pageSize).Offset((current - 1) * pageSize).Find(ctx)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	total, err := db.Count(ctx)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	resp := shop.CategoryListResp{
		Total: total,
	}
	copier.Copy(&resp.Data, &list)

	return &resp, nil
}
