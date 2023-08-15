package categoryservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/shop/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategorySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategorySaveLogic {
	return &CategorySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategorySaveLogic) updatePath(db model.ProductCategoryModel, id int64, parentId int64) (path string, err error) {
	ctx := l.ctx

	idStr := strconv.FormatInt(id, 10)

	parentIdPath := "0-"

	if parentId == 0 {
		path = parentIdPath + idStr
	} else {
		query := "parent_id = ?"
		queryArgs := []interface{}{parentId}
		pageSize := 1
		current := 10
		var parentModel *model.ProductCategory
		parentModel, err = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).First(ctx)
		if err != nil {
			return
		}

		if !util.IsStringEmpty(parentModel.Path) {
			parentIdPath = parentModel.Path + "-"
		}

		path = parentIdPath + idStr
	}

	return

}

func (l *CategorySaveLogic) CategorySave(in *shop.CategorySaveReq) (*shop.CategoryResp, error) {
	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	config := conf.Database.NewConf(in.GetSiteId())
	dsn := config.Dsn()

	//mysql model调用
	db := model.NewProductCategoryModel(sqlx.NewMysql(dsn), conf.Cache)

	id := in.GetId()

	goodsCategory := model.ProductCategory{}
	now := time.Now().Unix()
	// 如果不存在id则是新增

	var (
		err     error
		findOne *model.ProductCategory
		path    string
		insert  sql.Result
	)

	parentId := in.ParentId
	if parentId != nil && *parentId > 0 {
		_, err = db.FindOne(ctx, *parentId)
		if err != nil {
			return nil, err
		}
	}

	if id == 0 {
		err = copier.Copy(&goodsCategory, &in)
		if err != nil {
			return nil, err
		}
		goodsCategory.CreatedAt = now

		goodsCategory.UpdatedAt = now

		if in.Status == nil {
			goodsCategory.Status = 1
		}

		if in.ListOrder == nil {
			goodsCategory.ListOrder = 10000
		}

		insert, err = db.Insert(ctx, &goodsCategory)
		if err != nil {
			return nil, err
		}

		id, err = insert.LastInsertId()

		if err != nil {
			return nil, err
		}

		path, err = l.updatePath(db, id, goodsCategory.ParentId)
		if err != nil {
			return nil, err
		}

		goodsCategory.ProductCategoryId = id

		goodsCategory.Path = path

		err = db.Update(ctx, &goodsCategory)

	} else {

		// 查当前菜单是否存在
		findOne, err = db.FindOne(ctx, id)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, errors.New("该分类不存在！")
			}
			return nil, err
		}

		err = copier.Copy(&findOne, &in)
		if err != nil {
			return nil, err
		}

		goodsCategory = *findOne

		path, err = l.updatePath(db, id, goodsCategory.ParentId)

		goodsCategory.Path = path

		goodsCategory.UpdatedAt = now

		err = db.Update(ctx, &goodsCategory)
		if err != nil {
			return nil, err
		}

	}

	resp := shop.CategoryResp{}
	err = copier.Copy(&resp, &goodsCategory)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
