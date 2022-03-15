/**
** @创建时间: 2021/1/10 1:39 下午
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"errors"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"gorm.io/gorm"
)

type ThemeFile struct {
	controller.Rest
}

type option struct {
	Id   int    `json:"id"`
	File string `json:"file"`
	Name string `json:"name"`
}

func (rest ThemeFile) List(c *gin.Context) {
	t := c.Query("type")
	opt := model.Option{}
	db := util.GetDb(c)
	tx := db.Where("option_name", "theme").First(&opt)
	gorm.ErrRecordNotFound = errors.New("主题不能为空！")
	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}
	theme := opt.OptionValue
	var list []model.ThemeFile
	tx = db.Where("theme = ? AND type = ?", theme, t).Find(&list)
	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	file := "list"
	if t == "article" {
		file = "article"
	} else if t == "page" {
		file = "page"
	}

	result := []option{{Name: "默认模板", File: file}}

	for _, v := range list {
		result = append(result, option{
			Id:   v.Id,
			Name: v.Name,
			File: v.File,
		})
	}
	rest.Success(c, "获取成功！", result)
}

func (rest *ThemeFile) Save(c *gin.Context) {

	id := c.Param("id")
	var form struct {
		More string `json:"more"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	more := form.More
	if more == "" {
		rest.Error(c, "配置不能为空！", nil)
		return
	}

	db := util.GetDb(c)
	themeFile := model.ThemeFile{}
	tx := db.Where("id = ?", id).First(&themeFile)
	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	tx = db.Model(&model.ThemeFile{}).Where("id = ?", id).Update("more", more)
	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}
	rest.Success(c, "更新成功！", nil)
}
