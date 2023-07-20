package categoryservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/shop/model/modelgoods"

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

func (l *CategorySaveLogic) updatePath(db modelgoods.GoodsCategoryModel, id int64, parentId int64) (path string, err error) {
	ctx := l.ctx

	idStr := strconv.FormatInt(id, 10)

	parentIdPath := "0-"

	if parentId == 0 {
		path = parentIdPath + idStr
	} else {
		query := "parent_id = ?"
		queryArgs := []interface{}{parentId}
		var pageSize int32 = 1
		var current int32 = 10
		var parentModel *modelgoods.GoodsCategory
		parentModel, err = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).First(ctx)
		if err != nil {
			return
		}

		if !util.IsStringEmpty(parentModel.Path.String) {
			parentIdPath = parentModel.Path.String + "-"
		}

		path = parentIdPath + idStr
	}

	return

}

func (l *CategorySaveLogic) CategorySave(in *shop.CategorySaveReq) (*shop.CategoryResp, error) {
	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")

	//mysql model调用
	db := modelgoods.NewGoodsCategoryModel(sqlx.NewMysql(dsn), conf.Cache)

	id := in.GetId()

	goodsCategory := modelgoods.GoodsCategory{}
	now := time.Now().Unix()
	// 如果不存在id则是新增

	var (
		err     error
		findOne *modelgoods.GoodsCategory
		path    string
		insert  sql.Result
	)

	parentId := in.ParentId
	if parentId != 0 {
		_, err = db.FindOne(ctx, parentId)
		if err != nil {
			return nil, err
		}
	}

	if id == 0 {
		err = copier.Copy(&goodsCategory, &in)
		if err != nil {
			return nil, err
		}
		goodsCategory.CreatedAt = sql.NullInt64{
			Int64: now,
			Valid: true,
		}

		goodsCategory.UpdatedAt = sql.NullInt64{
			Int64: now,
			Valid: true,
		}

		insert, err = db.Insert(ctx, &goodsCategory)
		if err != nil {
			return nil, err
		}

		id, err = insert.LastInsertId()

		if err != nil {
			return nil, err
		}

		path, err = l.updatePath(db, id, goodsCategory.ParentId.Int64)
		if err != nil {
			return nil, err
		}

		goodsCategory.Id = id

		goodsCategory.Path = sql.NullString{
			String: path,
			Valid:  true,
		}

		err = db.Update(ctx, &goodsCategory)

	} else {

		// 查当前菜单是否存在
		findOne, err = db.FindOne(ctx, id)
		if err != nil {
			if errors.Is(err, modelgoods.ErrNotFound) {
				return nil, errors.New("该分类不存在！")
			}
			return nil, err
		}

		err = copier.Copy(&findOne, &in)
		if err != nil {
			return nil, err
		}

		goodsCategory = *findOne

		path, err = l.updatePath(db, id, goodsCategory.ParentId.Int64)

		goodsCategory.Path = sql.NullString{
			String: path,
			Valid:  true,
		}

		goodsCategory.UpdatedAt = sql.NullInt64{
			Int64: now,
			Valid: true,
		}

		err = db.Update(ctx, &goodsCategory)
		if err != nil {
			return nil, err
		}

	}

	resp := shop.CategoryResp{}
	copier.Copy(&resp, &goodsCategory)
	return &resp, nil
}
