package categoryservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/model"

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
	db := model.NewProductCategoryModel(sqlx.NewMysql(dsn), conf.Cache)

	current := in.GetCurrent()
	pageSize := in.GetPageSize()

	// 获取分类列表
	list, err := db.Limit(int(pageSize)).Offset(int((current - 1) * pageSize)).Find(ctx)
	if err != nil {
		return nil, err
	}

	var total int64
	if pageSize > 0 {
		total, err = db.Count(ctx)
		if err != nil {
			return nil, err
		}
	}

	var _data = make([]*shop.CategoryResp, len(list))
	for k, v := range list {
		item := shop.CategoryResp{}
		copier.Copy(&item, &v)
		if v.CreatedAt > 0 {
			createdAt := v.CreatedAt
			item.CreatedTime = time.Unix(createdAt, 0).Format(data.TimeLayout)
		}
		if v.UpdatedAt > 0 {
			updatedAt := v.UpdatedAt
			item.UpdatedTime = time.Unix(updatedAt, 0).Format(data.TimeLayout)
		}
		_data[k] = &item

	}

	resp := shop.CategoryListResp{
		Total: total,
		Data:  _data,
	}

	return &resp, nil
}
