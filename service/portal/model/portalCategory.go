/**
** @创建时间: 2020/10/29 4:47 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/util"
)

type PortalCategories struct {
	Id             int     `json:"id"`
	ParentId       int     `gorm:"type:bigint(20);comment:父级id;not null" json:"parent_id"`
	PostCount      int     `gorm:"type:bigint(20);comment:分类文章数;not null" json:"post_count"`
	Status         int     `gorm:"type:tinyint(3);comment:状态,1:发布,0:不发布;default:1;not null" json:"status"`
	DeleteAt       int64   `gorm:"type:bigint(20);comment:删除时间;not null" json:"delete_at"`
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
	TopAlias       string  `gorm:"-" json:"top_alias"`
}

type PortalTree struct {
	PortalCategories
	Children []PortalTree `json:"children"`
}

type CategoriesOptions struct {
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

func (model *PortalCategories) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model *PortalCategories) recursionByParent(Categories []PortalCategories, parentId int) []PortalTree {
	var tree []PortalTree
	for _, v := range Categories {
		// 当前子项
		if parentId == v.ParentId {
			item := PortalTree{
				PortalCategories: v,
			}

			children := model.recursionByParent(Categories, v.Id)
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

func (model *PortalCategories) Index(db *gorm.DB, query string, queryArgs []interface{}) (data []PortalTree, err error) {
	// 获取默认的系统分页

	var Categories []PortalCategories
	tx := db.Where(query, queryArgs...).Find(&Categories)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return
		}
	}
	// 生成树形结构
	data = model.recursionChildById(Categories, 0)
	return data, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 显示全部门户分类
 * @Date 2022/2/11 8:50:44
 * @Param
 * @return
 **/

func (model *PortalCategories) List(db *gorm.DB) ([]PortalCategories, error) {
	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}
	queryStr := strings.Join(query, " AND ")
	var Categories []PortalCategories
	tx := db.Where(queryStr, queryArgs...).Find(&Categories)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return Categories, tx.Error
		}
	}

	for k, v := range Categories {

		topCategories := new(PortalCategories)
		err := topCategories.GetTopCategories(db, v.Id)
		if err == nil {
			Categories[k].TopAlias = topCategories.Alias
			Categories[k].PrevPath = util.FileUrl(v.Thumbnail)
		}

	}

	return Categories, nil
}

func (model *PortalCategories) recursionParent(Categories []PortalCategories, id int) (topId int) {
	topId = id
	for _, v := range Categories {
		if v.Id == id && v.ParentId > 0 {
			topId = model.recursionParent(Categories, v.ParentId)
		}
	}
	return topId
}

func (model *PortalCategories) indent(level int) string {

	indent := ""
	for i := 0; i < level; i++ {
		indent += "    |--"
	}

	return indent

}

func (model *PortalCategories) GetTopId(db *gorm.DB, id int) (int, error) {
	tx := db.Where("id = ? AND delete_at = ?", id, 0).First(&model)
	if tx.Error != nil {
		return 0, tx.Error
	}
	var Categories []PortalCategories
	tx = db.Where("delete_at = ?", 0).Find(&Categories)
	if tx.Error != nil {
		return 0, tx.Error
	}
	topId := model.recursionParent(Categories, id)
	return topId, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取分类别名路径
 * @Date 2022/2/11 8:51:44
 * @Param
 * @return
 **/

func (model *PortalCategories) GetAlias() (url string) {
	alias := "/" + model.Alias
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

func (model *PortalCategories) GetPrevious(db *gorm.DB, id int) (breadcrumbs []Breadcrumb, err error) {
	tx := db.Where("id = ? AND delete_at = ?", id, 0).First(&model)
	if tx.Error != nil {
		return breadcrumbs, tx.Error
	}
	var Categories []PortalCategories
	tx = db.Where("delete_at = ?", 0).Order("parent_id asc").Find(&Categories)
	if tx.Error != nil {
		return breadcrumbs, tx.Error
	}

	breadcrumbs = model.recursionPrevious(Categories, model.ParentId)

	breadcrumbs = append(breadcrumbs, Breadcrumb{
		Id:    model.Id,
		Name:  model.Name,
		Alias: model.Alias,
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

func (model *PortalCategories) recursionPrevious(Categories []PortalCategories, parentId int) (breadcrumbs []Breadcrumb) {
	for _, v := range Categories {
		if v.Id == parentId {
			breadcrumbs = append(breadcrumbs, Breadcrumb{
				Id:    v.Id,
				Name:  v.Name,
				Alias: v.Alias,
			})
			childBreadcrumbs := model.recursionPrevious(Categories, v.ParentId)
			breadcrumbs = append(childBreadcrumbs, breadcrumbs...)
		}
	}
	return breadcrumbs
}

func (model *PortalCategories) Show(db *gorm.DB, query string, queryArgs []interface{}) (err error) {
	result := db.Where(query, queryArgs...).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = errors.New("该分类不存在！")
			return
		}
		err = result.Error
		return
	}
	if model.Thumbnail != "" {
		model.PrevPath = util.FileUrl(model.Thumbnail)
	}
	//Categories.Alias = Categories.GetAlias()
	return
}

func (model *PortalCategories) Save(db *gorm.DB) (PortalCategories, error) {

	var tx *gorm.DB
	if model.Id == 0 {
		tx = db.Create(&model)
	}

	tx = db.Save(&model)
	if tx.Error != nil {
		return PortalCategories{}, tx.Error
	}
	return *model, nil
}

func (model *PortalCategories) GetTopCategories(db *gorm.DB, id int) (err error) {
	var cid int
	cid, err = model.GetTopId(db, id)
	if err != nil {
		return
	}

	err = model.Show(db, "id = ? AND delete_at = ?", []interface{}{cid, 0})
	if err != nil {
		return
	}
	return
}

func (model *PortalCategoriesPost) Store(db *gorm.DB, pcpPost []PortalCategoriesPost) ([]PortalCategoriesPost, error) {

	var pcp []PortalCategoriesPost
	result := db.Where("post_id  = ?", model.PostId).Find(&pcp)
	if result.Error != nil {
		return pcp, nil
	}

	// 删除原来的
	var delQuery []string
	var delQueryArgs []interface{}

	for _, v := range pcp {
		if !model.inArray(v, pcpPost) || len(pcpPost) == 0 {
			delQuery = append(delQuery, "(post_id = ? and Categories_id = ?)")
			delQueryArgs = append(delQueryArgs, v.PostId, v.CategoriesId)
		}

		// 如果未传参，全部删除
		if len(pcpPost) == 0 {
			delQuery = append(delQuery, "(post_id = ? and Categories_id = ?)")
			delQueryArgs = append(delQueryArgs, v.PostId, v.CategoriesId)
		}
	}

	var toAddPcp []PortalCategoriesPost

	// 添加待添加的
	for _, v := range pcpPost {
		if !model.inArray(v, pcp) || len(pcp) == 0 {
			toAddPcp = append(toAddPcp, PortalCategoriesPost{
				PostId:       v.PostId,
				CategoriesId: v.CategoriesId,
			})
		}
	}

	// 删除要删除的
	delQueryStr := strings.Join(delQuery, " OR ")
	if delQueryStr != "" {
		db.Where(delQueryStr, delQueryArgs...).Delete(&PortalCategoriesPost{})
	}

	//添加要添加的
	if len(toAddPcp) > 0 {
		result = db.Create(&toAddPcp)
		if result.Error != nil {
			return []PortalCategoriesPost{}, nil
		}
	}

	// 查询最后的结果
	result = db.Where("post_id  = ?", model.Id).Find(&pcp)
	if result.Error != nil {
		return pcp, nil
	}
	return pcp, nil
}

func (model *PortalCategoriesPost) inArray(inPost PortalCategoriesPost, pcp []PortalCategoriesPost) bool {

	for _, v := range pcp {

		if inPost.PostId == v.PostId && inPost.CategoriesId == v.CategoriesId {
			return true
		}
	}
	return false
}

func (model *PortalCategories) ListWithTree(db *gorm.DB) ([]PortalTree, error) {

	tree, err := model.List(db)
	if err != nil {
		return []PortalTree{}, err
	}

	// 生成树形结构
	data := model.recursionChildById(tree, model.ParentId)

	if len(data) == 0 {
		data = make([]PortalTree, 0)
	}

	return data, nil
}

func (model *PortalCategories) recursionChildById(Categories []PortalCategories, parentId int) []PortalTree {

	var tree []PortalTree

	for _, v := range Categories {

		// 当前子项
		if parentId == v.ParentId {

			item := PortalTree{
				PortalCategories: v,
			}

			if parentId == 0 || v.ParentId > 0 {
				children := model.recursionChildById(Categories, v.Id)
				item.Children = children
			}

			tree = append(tree, item)

		}
	}

	return tree

}

func (model *PortalCategories) ListWithOptions(db *gorm.DB, query string, queryArgs []interface{}) ([]CategoriesOptions, error) {
	var pc []PortalCategories
	cOptions := make([]CategoriesOptions, 0)
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

func (model *PortalCategories) recursionOptions(nav []PortalCategories, parentId int, level int) (cOptions []CategoriesOptions) {
	nextLevel := level + 1
	for _, v := range nav {
		if parentId == v.ParentId {
			ops := CategoriesOptions{
				Id:       v.Id,
				ParentId: v.ParentId,
				Name:     v.Name,
				Level:    level,
			}
			cOptions = append(cOptions, ops)
			childs := model.recursionOptions(nav, v.Id, nextLevel)
			cOptions = append(cOptions, childs...)
		}
	}
	return cOptions
}

// 获取子集的分类id

func (model *PortalCategories) ChildIds(db *gorm.DB, id int) ([]string, error) {
	tree, err := model.List(db)
	if err != nil {
		return []string{}, nil
	}
	ids := model.recursionChild(tree, id)
	ids = append(ids, strconv.Itoa(id))
	return ids, nil
}

func (model *PortalCategories) recursionChild(Categories []PortalCategories, parentId int) []string {
	var ids []string
	for _, v := range Categories {
		if parentId == v.ParentId {
			ids = append(ids, strconv.Itoa(v.Id))
			childIds := model.recursionChild(Categories, v.Id)

			ids = append(ids, childIds...)
		}
	}
	return ids
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除一项
 * @Date 2020/11/8 19:27:07
 * @Param
 * @return
 **/

func (model *PortalCategories) Delete(db *gorm.DB) (err error) {
	id := model.Id
	if id == 0 {
		err = errors.New("分类id不能为0或空！")
		return
	}

	portalCategories := new(PortalCategories)

	err = portalCategories.Show(db, "id = ? and delete_at = ?", []interface{}{id, 0})
	if err != nil {
		return
	}

	// 查看当前分类下是否存在子分类

	var count int64
	db.Model(model).Where("parent_id = ? AND delete_at = ?", id, 0).Count(&count)

	if count > 0 {
		err = errors.New("请先删除分类下的子分类！")
		return
	}

	deleteAt := time.Now().Unix()
	tx := db.Model(model).Where("id = ?", id).Update("delete_at", deleteAt)

	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 批量删除
 * @Date 2020/11/8 19:41:45
 * @Param
 * @return
 **/

func (model *PortalCategories) BatchDelete(db *gorm.DB, ids []string) (err error) {
	deleteAt := time.Now().Unix()
	if err = db.Model(&model).Where("id IN (?)", ids).Updates(map[string]interface{}{"delete_at": deleteAt}).Error; err != nil {
		return
	}
	return
}
