package categoryservicelogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/model"

	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryTreeLogic {
	return &CategoryTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func buildTree(categories []*model.ProductCategory, parentId int64) []*shop.CategoryTreeData {
	var treeData []*shop.CategoryTreeData
	for _, category := range categories {
		if category.ParentId.Valid && category.ParentId.Int64 == parentId {
			item := shop.CategoryTreeData{}
			copier.Copy(&item, &category)
			if category.CreatedAt > 0 {
				createdAt := category.CreatedAt
				item.CreatedTime = time.Unix(createdAt, 0).Format(data.TimeLayout)
			}
			if category.UpdatedAt > 0 {
				updatedAt := category.UpdatedAt
				item.UpdatedTime = time.Unix(updatedAt, 0).Format(data.TimeLayout)
			}
			item.Children = buildTree(categories, category.ProductCategoryId)
			treeData = append(treeData, &item)
		}
	}

	return treeData
}
func (l *CategoryTreeLogic) CategoryTree(in *shop.CategoryTreeReq) (*shop.CategoryTreeListResp, error) {

	ctx := l.ctx
	c := l.svcCtx
	conf := c.Config
	dsn := conf.Database.Dsn("")
	//mysql model调用
	db := model.NewProductCategoryModel(sqlx.NewMysql(dsn), conf.Cache)

	ignoreId := in.IgnoreId
	name := in.Name
	status := in.Status

	var query = make([]string, 0)
	var queryArgs = make([]interface{}, 0)
	if ignoreId != nil && *ignoreId > 0 {
		query = append(query, "product_category_id != ?")
		queryArgs = append(queryArgs, ignoreId)
	}

	if strings.TrimSpace(name) != "" {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	if status != nil {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, status)
	}

	queryStr := strings.Join(query, " and ")
	// 获取分类列表
	list, err := db.Where(queryStr, queryArgs...).Find(ctx)
	if err != nil {
		return nil, err
	}

	treeData := buildTree(list, 0)

	return &shop.CategoryTreeListResp{
		Data: treeData,
	}, nil
}
