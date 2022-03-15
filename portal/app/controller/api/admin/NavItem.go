/**
** @创建时间: 2021/1/4 9:32 上午
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"errors"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"gorm.io/gorm"
	"strconv"
)

type NavItem struct {
	controller.Rest
}

type Options struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionsMap struct {
	Label   string    `json:"label"`
	Options []Options `json:"options"`
}

func (rest *NavItem) Get(c *gin.Context) {

	navId := c.Query("nav_id")

	if navId == "" {
		rest.Error(c, "导航不能为空！", nil)
		return
	}

	var query = "nav_id  = ?"
	var queryArgs = []interface{}{navId}

	db := util.GetDb(c)

	navItem, err := new(model.NavItem).GetWithChild(db, query, queryArgs)

	if util.IsDbErr(db) != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", navItem)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 根据nav_key获取全部导航列表
 * @Date 2022/1/22 13:22:28
 * @Param
 * @return
 **/

func (rest *NavItem) GetNavItems(c *gin.Context) {

	var form struct {
		Key string `json:"key"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	key := form.Key

	if key == "" {
		rest.Error(c, "唯一标识不能为空！", nil)
		return
	}

	db := util.GetDb(c)
	query := "`key` = ?"
	queryArgs := []interface{}{key}
	nav := model.Nav{}
	err := nav.Show(db, query, queryArgs)
	if err != nil {
		rest.Error(c, "操作失败", nil)
		return
	}

	if nav.Id == 0 {
		nav.Key = key
		tx := db.Where(query, queryArgs...).FirstOrCreate(&nav)
		if util.IsDbErr(tx) != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}
	}

	// 根据navId获取全部导航项
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	itemQuery := "nav_id = ? AND status = 1"
	itemQueryArgs := []interface{}{nav.Id}

	navItemsPaginate, err := new(model.NavItem).GetWithChildPaginate(db, current, pageSize, itemQuery, itemQueryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	navItems, err := new(model.NavItem).GetWithChild(db, itemQuery, itemQueryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "操作成功！", gin.H{
		"navId":    nav.Id,
		"navItemsPaginate": navItemsPaginate,
		"navItems":navItems,
	})

}

func (rest *NavItem) OptionsList(c *gin.Context) {

	navId := c.Query("nav_id")

	if navId == "" {
		rest.Error(c, "导航不能为空！", nil)
		return
	}

	var query = "nav_id  = ?"
	var queryArgs = []interface{}{navId}

	db := util.GetDb(c)

	result := new(model.NavItem).OptionsList(db, query, queryArgs)

	rest.Success(c, "获取成功！", result)
	return

}

func (rest *NavItem) OptionsUrls(c *gin.Context) {

	db := util.GetDb(c)

	portalCategory, err := new(model.PortalCategory).List(db)

	if err != nil {
		rest.Error(c, err.Error(), nil)
	}

	categoryOptions := make([]Options, 0)

	for _, v := range portalCategory {

		var url = "/list/" + strconv.Itoa(v.Id)
		if v.Alias != "" {
			url = "/" + v.Alias
		}

		categoryOptions = append(categoryOptions, Options{
			Label: v.Name,
			Value: url,
		})

	}

	pages, err := model.PortalPost{PostType: 2}.PortalList(db, "", nil)

	pageOptions := make([]Options, 0)

	for _, v := range pages {
		pageOptions = append(pageOptions, Options{
			Label: v.PostTitle,
			Value: "/page/" + strconv.Itoa(v.Id),
		})

	}

	var om = []OptionsMap{{
		Label: "首页",
		Options: []Options{
			{
				Label: "首页",
				Value: "/",
			}},
	}, {
		Label:   "文章分类",
		Options: categoryOptions,
	}, {
		Label:   "所有页面",
		Options: pageOptions,
	}}

	rest.Success(c, "获取成功！", om)

}

func (rest *NavItem) Show(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	db := util.GetDb(c)

	data, err := new(model.NavItem).Show(db, "id = ?", []interface{}{rewrite.Id})

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", data)
}

func (rest *NavItem) Edit(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	editId := rewrite.Id
	rest.Save(c, editId)
}

func (rest *NavItem) Store(c *gin.Context) {
	rest.Save(c, 0)
}

// 保存新增编辑的内容

func (rest *NavItem) Save(c *gin.Context, editId int) {

	var form struct {
		NavId     int     `json:"nav_id"`
		ParentId  int     `json:"parent_id"`
		Status    int     `json:"status"`
		ListOrder float64 `json:"list_order"`
		Name      string  `json:"name"`
		Target    string  `json:"target"`
		HrefType  int     `json:"href_type"`
		Href      string  `json:"href"`
		Icon      string  `json:"icon"`
		Path      string  `json:"path"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if form.NavId == 0 {
		rest.Error(c, "导航id不能为空！", nil)
		return
	}

	if form.Name == "" {
		rest.Error(c, "导航项名称不能为空！", nil)
		return
	}

	navItem := model.NavItem{
		NavId:     form.NavId,
		ParentId:  form.ParentId,
		Status:    form.Status,
		ListOrder: form.ListOrder,
		Name:      form.Name,
		Target:    form.Target,
		Href:      form.Href,
		Icon:      form.Icon,
		Path:      form.Path,
	}

	db := util.GetDb(c)

	msg := ""
	if editId > 0 {
		tempNavItem := model.NavItem{}
		tx := db.Where("id = ?", editId).First(&tempNavItem)

		if tx.Error != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		navItem.Id = tempNavItem.Id
		db.Save(&navItem)
		msg = "保存成功！"
	} else {
		db.Create(&navItem)
		msg = "创建成功！"
	}

	rest.Success(c, msg, navItem)

}

func (rest *NavItem) Delete(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	db := util.GetDb(c)

	// 查询是否存在子分类
	var navItem []model.NavItem
	tx := db.Where("parent_id = ?", rewrite.Id).Find(&navItem)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	if tx.RowsAffected > 0 {
		rest.Error(c, "请先删除子分类！", nil)
		return
	}

	tx = db.Where("id = ?", rewrite.Id).Delete(&navItem)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, "删除成功！", nil)
}
