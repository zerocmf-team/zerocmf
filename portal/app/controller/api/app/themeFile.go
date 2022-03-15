/**
** @创建时间: 2021/1/7 3:08 下午
** @作者　　: return
** @描述　　:
 */

package app

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

func (rest *ThemeFile) List(c *gin.Context) {

	theme := c.Query("theme")

	if theme == "" {
		rest.Error(c,"主题不能为空！",nil)
		return
	}

	isPublic := c.Query("is_public")

	query := "theme = ? AND is_public = ?"
	queryArgs := []interface{}{theme, isPublic}

	db := util.GetDb(c)

	data,err := new(model.ThemeFile).List(db,query,queryArgs)

	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound) {
		rest.Error(c,err.Error(),nil)
		return
	}

	rest.Success(c,"获取成功！",data)


}

func (rest *ThemeFile) Detail(c *gin.Context) {

	theme := c.Query("theme")

	if theme == "" {
		rest.Error(c,"主题不能为空！",nil)
		return
	}

	file := c.Query("file")
	if file == "" {
		rest.Error(c,"文件不能为空！",nil)
		return
	}

	query := "theme = ? AND file = ?"
	queryArgs := []interface{}{theme, file}

	db := util.GetDb(c)

	data,err := new(model.ThemeFile).Show(db,query,queryArgs)

	if err != nil {
		rest.Error(c,err.Error(),nil)
		return
	}

	rest.Success(c,"获取成功！",data)

}
