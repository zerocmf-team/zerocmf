package categoryservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zerocmf/service/shop/model"

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
	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")
	//mysql model调用
	db := model.NewProductCategoryModel(sqlx.NewMysql(dsn), conf.Cache)
	id := in.GetId()
	goodCategory, err := db.FindOne(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("该内容不存在！")
		}
		return nil, err
	}
	goodCategory.DeletedAt = time.Now().Unix()
	err = db.Update(ctx, goodCategory)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
	}
	resp := shop.CategoryResp{}
	copier.Copy(&resp, &goodCategory)
	return &resp, nil
}
