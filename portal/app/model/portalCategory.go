/**
** @创建时间: 2020/10/29 4:47 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type PortalCategory struct {
	Id             int     `json:"id"`
	ParentId       int     `gorm:"type:bigint(20);comment:父级id;not null" json:"parent_id"`
	PostCount      int     `gorm:"type:bigint(20);comment:分类文章数;not null" json:"post_count"`
	Status         int     `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;default:1;not null" json:"status"`
	DeleteAt       int64   `gorm:"type:int(11);comment:删除时间;not null" json:"delete_at"`
	ListOrder      float64 `gorm:"type:float(0);comment:排序;default:10000;not null" json:"list_order"`
	Name           string  `gorm:"type:varchar(200);comment:唯一名称;not null" json:"name"`
	Alias          string  `gorm:"type:varchar(200);comment:唯一名称;not null" json:"alias"`
	Description    string  `gorm:"type:varchar(255);comment:分类描述;not null" json:"description"`
	Thumbnail      string  `gorm:"type:varchar(255);comment:缩略图;not null" json:"thumbnail"`
	Path           string  `gorm:"type:varchar(255);comment:分类层级关系;not null" json:"path"`
	SeoTitle       string  `gorm:"type:varchar(100);comment:三要素标题;not null" json:"seo_title"`
	SeoKeywords    string  `gorm:"type:varchar(255);comment:三要素关键字;not null" json:"seo_keywords"`
	SeoDescription string  `gorm:"type:varchar(255);comment:三要素描述;not null" json:"seo_description"`
	ListTpl        string  `gorm:"type:varchar(50);comment:分类列表模板;not null" json:"list_tpl"`
	OneTpl         string  `gorm:"type:varchar(50);comment:分类文章页模板;not null" json:"one_tpl"`
	More           string  `gorm:"type:longtext;comment:扩展属性" json:"more"`
	PrevPath       string  `gorm:"-" json:"prev_path"`
	TopSlug        string  `gorm:"-" json:"top_slug"`
}

type portalTree struct {
	PortalCategory
	Value    string       `json:"value"`
	Title    string       `json:"title"`
	Children []portalTree `json:"children"`
}

type categoryOptions struct {
	Id       int    `json:"id"`
	ParentId int    `gorm:"type:int(11);comment:所属父类id;default:0" json:"parent_id"`
	Name     string `gorm:"type:varchar(50);comment:路由名称" json:"name"`
	Level    int    `json:"level"`
}

type Breadcrumb struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

func (model PortalCategory) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model PortalCategory) recursionByParent(category []PortalCategory, parentId int) []portalTree {
	var tree []portalTree
	for _, v := range category {
		// 当前子项
		if parentId == v.ParentId {
			item := portalTree{
				PortalCategory: v,
				Value:          strconv.Itoa(v.Id),
				Title:          v.Name,
			}

			children := model.recursionByParent(category, v.Id)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归显示树形菜单
 * @Date 2022/2/11 8:49:59
 * @Param
 * @return
 **/

func (model *PortalCategory) Index(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (paginate.Paginate, error) {

	// 获取默认的系统分页
	// 合并参数合计
	var total int64 = 0
	var category []PortalCategory
	db.Where(query, queryArgs...).Find(&category).Count(&total)
	tx := db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&category)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return paginate.Paginate{}, tx.Error
		}
	}
	// 生成树形结构
	data := model.recursionByParent(category, 0)
	paginate := paginate.Paginate{Data: data, Current: current, PageSize: pageSize, Total: total}
	if len(category) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 显示全部门户分类
 * @Date 2022/2/11 8:50:44
 * @Param
 * @return
 **/

func (model *PortalCategory) List(db *gorm.DB) ([]PortalCategory, error) {
	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}
	queryStr := strings.Join(query, " AND ")
	var category []PortalCategory
	tx := db.Where(queryStr, queryArgs...).Find(&category)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return category, tx.Error
		}
	}

	for k, v := range category {
		topCategory, _ := new(PortalCategory).GetTopCategory(db, v.Id)
		category[k].TopSlug = topCategory.Alias
		category[k].PrevPath = util.FileUrl(v.Thumbnail)
	}

	return category, nil
}

func (model *PortalCategory) recursionParent(category []PortalCategory, id int) (topId int) {
	topId = id
	for _, v := range category {
		if v.Id == id && v.ParentId > 0 {
			topId = model.recursionParent(category, v.ParentId)
		}
	}
	return topId
}

func (model *PortalCategory) indent(level int) string {

	indent := ""
	for i := 0; i < level; i++ {
		indent += "    |--"
	}

	return indent

}

func (model *PortalCategory) GetTopId(db *gorm.DB, id int) (int, error) {
	tx := db.Where("id = ? AND delete_at = ?", id, 0).First(&model)
	if tx.Error != nil {
		return 0, tx.Error
	}
	var category []PortalCategory
	tx = db.Where("delete_at = ?", 0).Find(&category)
	if tx.Error != nil {
		return 0, tx.Error
	}
	topId := model.recursionParent(category, id)
	return topId, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取分类别名路径
 * @Date 2022/2/11 8:51:44
 * @Param
 * @return
 **/

func (model PortalCategory) GetAlias() (url string) {
	alias := "/"+model.Alias
	if model.Alias == "" {
		alias = "/list/" + strconv.Itoa(model.Id)
	}
	return alias
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 根据id递归获取全部上级的数据
 * @Date 2022/2/10 9:47:18
 * @Param
 * @return
 **/

func (model *PortalCategory) GetPrevious(db *gorm.DB, id int) (breadcrumbs []Breadcrumb, err error) {
	tx := db.Where("id = ? AND delete_at = ?", id, 0).First(&model)
	if tx.Error != nil {
		return breadcrumbs, tx.Error
	}
	var category []PortalCategory
	tx = db.Where("delete_at = ?", 0).Order("parent_id asc").Find(&category)
	if tx.Error != nil {
		return breadcrumbs, tx.Error
	}

	breadcrumbs = model.recursionPrevious(category, model.ParentId)

	breadcrumbs = append(breadcrumbs, Breadcrumb{
		Id:    model.Id,
		Name:  model.Name,
		Alias: model.GetAlias(),
	})

	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归所有上级的分类数据
 * @Date 2022/2/10 9:54:32
 * @Param
 * @return
 **/

func (model *PortalCategory) recursionPrevious(category []PortalCategory, parentId int) (breadcrumbs []Breadcrumb) {
	for _, v := range category {
		if v.Id == parentId {
			breadcrumbs = append(breadcrumbs, Breadcrumb{
				Id:    v.Id,
				Name:  v.Name,
				Alias: v.GetAlias(),
			})
			childBreadcrumbs := model.recursionPrevious(category, v.ParentId)
			breadcrumbs = append(childBreadcrumbs, breadcrumbs...)
		}
	}
	return breadcrumbs
}

func (model *PortalCategory) Show(db *gorm.DB, query string, queryArgs []interface{}) (PortalCategory, error) {
	category := PortalCategory{}
	result := db.Where(query, queryArgs...).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return category, errors.New("该分类不存在！")
		}
		return category, result.Error
	}
	if category.Thumbnail != "" {
		category.PrevPath = util.FileUrl(category.Thumbnail)
	}

	category.Alias = category.GetAlias()

	return category, nil
}

func (model *PortalCategory) Save(db *gorm.DB) (PortalCategory, error) {

	var tx *gorm.DB
	if model.Id == 0 {
		tx = db.Create(&model)
	}

	tx = db.Save(&model)
	if tx.Error != nil {
		return PortalCategory{}, tx.Error
	}
	return *model, nil
}

func (model *PortalCategory) GetTopCategory(db *gorm.DB, id int) (PortalCategory, error) {
	cid, err := model.GetTopId(db, id)
	if err != nil {
		return PortalCategory{}, err
	}
	data, err := model.Show(db, "id = ? AND delete_at = ?", []interface{}{cid, 0})
	if err != nil {
		return PortalCategory{}, err
	}
	return data, nil
}

func (model PortalCategoryPost) Store(db *gorm.DB, pcpPost []PortalCategoryPost) ([]PortalCategoryPost, error) {

	var pcp []PortalCategoryPost
	result := db.Where("post_id  = ?", model.PostId).Find(&pcp)
	if result.Error != nil {
		return pcp, nil
	}

	// 删除原来的
	var delQuery []string
	var delQueryArgs []interface{}

	for _, v := range pcp {
		if !model.inArray(v, pcpPost) || len(pcpPost) == 0 {
			delQuery = append(delQuery, "(post_id = ? and category_id = ?)")
			delQueryArgs = append(delQueryArgs, v.PostId, v.CategoryId)
		}

		// 如果未传参，全部删除
		if len(pcpPost) == 0 {
			delQuery = append(delQuery, "(post_id = ? and category_id = ?)")
			delQueryArgs = append(delQueryArgs, v.PostId, v.CategoryId)
		}
	}

	var toAddPcp []PortalCategoryPost

	// 添加待添加的
	for _, v := range pcpPost {
		if !model.inArray(v, pcp) || len(pcp) == 0 {
			toAddPcp = append(toAddPcp, PortalCategoryPost{
				PostId:     v.PostId,
				CategoryId: v.CategoryId,
			})
		}
	}

	// 删除要删除的
	delQueryStr := strings.Join(delQuery, " OR ")
	if delQueryStr != "" {
		db.Where(delQueryStr, delQueryArgs...).Delete(&PortalCategoryPost{})
	}

	//添加要添加的
	if len(toAddPcp) > 0 {
		result = db.Create(&toAddPcp)
		if result.Error != nil {
			return []PortalCategoryPost{}, nil
		}
	}

	// 查询最后的结果
	result = db.Where("post_id  = ?", model.Id).Find(&pcp)
	if result.Error != nil {
		return pcp, nil
	}
	return pcp, nil
}

func (model PortalCategoryPost) inArray(inPost PortalCategoryPost, pcp []PortalCategoryPost) bool {

	for _, v := range pcp {

		if inPost.PostId == v.PostId && inPost.CategoryId == v.CategoryId {
			return true
		}
	}
	return false
}

func (model *PortalCategory) ListWithTree(db *gorm.DB) ([]portalTree, error) {

	tree, err := model.List(db)
	if err != nil {
		return []portalTree{}, err
	}

	// 生成树形结构
	data := model.recursionChildById(tree, model.ParentId)

	if len(data) == 0 {
		data = make([]portalTree, 0)
	}

	return data, nil
}

func (model PortalCategory) recursionChildById(category []PortalCategory, parentId int) []portalTree {

	var tree []portalTree

	for _, v := range category {

		// 当前子项
		if parentId == v.ParentId {

			v.Alias = v.GetAlias()

			item := portalTree{
				PortalCategory: v,
				Value:          strconv.Itoa(v.Id),
				Title:          v.Name,
			}

			if parentId == 0 || v.ParentId > 0 {
				children := model.recursionChildById(category, v.Id)
				item.Children = children
			}

			tree = append(tree, item)

		}
	}

	return tree

}


func (model *PortalCategory) ListWithOptions(db *gorm.DB, query string, queryArgs []interface{}) ([]categoryOptions, error) {
	var pc []PortalCategory
	cOptions := make([]categoryOptions, 0)
	tx := db.Where(query, queryArgs...).Find(&pc)
	if tx.Error != nil {
		return cOptions, nil
	}
	data := model.recursionOptions(pc, 0, 0)
	for k, v := range data {
		data[k].Name = model.indent(v.Level) + v.Name
	}
	return data, nil
}

func (model *PortalCategory) recursionOptions(nav []PortalCategory, parentId int, level int) (cOptions []categoryOptions) {
	nextLevel := level + 1
	for _, v := range nav {
		if parentId == v.ParentId {
			ops := categoryOptions{
				Id:       v.Id,
				ParentId: v.ParentId,
				Name:     v.Name,
				Level:    level,
			}
			cOptions = append(cOptions, ops)
			childs := model.recursionOptions(nav, v.Id, nextLevel)
			cOptions = append(cOptions,childs...)
		}
	}
	return cOptions
}

// 获取子集的分类id

func (model *PortalCategory) ChildIds(db *gorm.DB, id int) ([]string, error) {
	tree, err := model.List(db)
	if err != nil {
		return []string{}, nil
	}
	ids := model.recursionChild(tree, id)
	ids = append(ids, strconv.Itoa(id))
	return ids, nil
}

func (model PortalCategory) recursionChild(category []PortalCategory, parentId int) []string {
	var ids []string
	for _, v := range category {
		if parentId == v.ParentId {
			ids = append(ids, strconv.Itoa(v.Id))
			childIds := model.recursionChild(category, v.Id)

			ids = append(ids, childIds...)
		}
	}
	return ids
}
