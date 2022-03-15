/**
** @创建时间: 2021/12/13 19:16
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	cmfModel "github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
	"strconv"
	"strings"
)

type PortalCategory struct {
	controller.Rest
}

func (rest *PortalCategory) Get(c *gin.Context) {

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	name := c.Query("name")
	if name != "" {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	queryStr := strings.Join(query, " AND ")

	db := util.GetDb(c)

	current, pageSize, err := new(paginate.Paginate).Default(c)

	data, err := new(model.PortalCategory).Index(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", data)
}

func (rest *PortalCategory) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	db := util.GetDb(c)
	data, err := new(model.PortalCategory).Show(db, "id = ? and delete_at = ?", []interface{}{rewrite.Id, 0})
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

func (rest *PortalCategory) List(c *gin.Context) {
	category := model.PortalCategory{
		ParentId: 0,
	}
	db := util.GetDb(c)
	data, err := category.ListWithTree(db)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

func (rest *PortalCategory) Store(c *gin.Context) {
	rest.Save(c, "0")
}

func (rest *PortalCategory) Edit(c *gin.Context) {
	id := c.Param("id")
	rest.Save(c, id)
}

func (rest *PortalCategory) Save(c *gin.Context, editId string) {

	var form struct {
		Name           string `json:"name" binding:"required"`
		ParentId       string `json:"parent_id"`
		Status         int    `json:"status"`
		Alias          string `json:"alias"`
		Description    string `json:"description"`
		Thumbnail      string `json:"thumbnail"`
		SeoTitle       string `json:"seo_title"`
		SeoDescription string `json:"seo_description"`
		SeoKeywords    string `json:"seo_keywords"`
		ListTpl        string `json:"list_tpl"`
		OneTpl         string `json:"one_tpl"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	db := util.GetDb(c)
	name := form.Name

	parentId, err := strconv.Atoi(form.ParentId)
	if err != nil {
		rest.Error(c, "parentId错误！", err.Error())
		return
	}
	portalCategory := model.PortalCategory{
		ParentId: parentId,
		Name:     name,
	}

	msg := "新增成功！"
	if editId != "" {
		id, err := strconv.Atoi(editId)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}
		portalCategory.Id = id
		msg = "更新成功！"
	}

	status := form.Status
	portalCategory.Status = status

	description := form.Description
	if description != "" {
		portalCategory.Description = description
	}
	thumbnail := form.Thumbnail
	if thumbnail != "" {
		portalCategory.Thumbnail = thumbnail
	}
	seoTitle := form.SeoTitle
	if seoTitle != "" {
		portalCategory.SeoTitle = seoTitle
	}
	seoDescription := form.SeoDescription
	if seoDescription != "" {
		portalCategory.SeoDescription = seoDescription
	}
	seoKeywords := form.SeoKeywords
	if seoKeywords != "" {
		portalCategory.SeoKeywords = seoKeywords
	}
	listTpl := form.ListTpl
	if listTpl != "" {
		portalCategory.ListTpl = listTpl
	}
	oneTpl := form.OneTpl
	if oneTpl != "" {
		portalCategory.OneTpl = oneTpl
	}

	alias := form.Alias

	if alias != "" {
		portalCategory.Alias = alias
	}

	tx, err := portalCategory.Save(db)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if alias != "" {
		fullUrl := "list/" + strconv.Itoa(portalCategory.Id)
		url := alias
		route := cmfModel.Route{
			Type:         1,
			FullUrl:      fullUrl,
			Url:          url,
		}
		err := route.Set(db)
		if err != nil {
			rest.Error(c,err.Error(),nil)
			return
		}
	}

	rest.Success(c, msg, tx)
}

func (rest *PortalCategory) Options(c *gin.Context) {

	category := model.PortalCategory{}
	var query = []string{"delete_at  = ?"}
	var queryArgs = []interface{}{0}
	queryStr := strings.Join(query, " AND ")
	db := util.GetDb(c)
	data, err := category.ListWithOptions(db, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", data)
}
