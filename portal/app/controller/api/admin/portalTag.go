/**
** @创建时间: 2021/2/6 10:32 上午
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
	"strings"
)

type Tag struct {
	controller.Rest
}

func (rest *Tag) Get(c *gin.Context) {
	name := c.Query("name")
	var query []string
	var queryArgs []interface{}
	if name != "" {
		query = []string{"name like ?"}
		queryArgs = []interface{}{"%" + name + "%"}
	}
	db := util.GetDb(c)
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	queryStr := strings.Join(query, " AND ")
	data, err := new(model.PortalTag).Index(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

func (rest *Tag) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	db := util.GetDb(c)
	var tag model.PortalTag
	tx := db.Where("id", rewrite.Id).First(&tag)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}
	if tx.RowsAffected == 0 {
		rest.Error(c, "内容不存在！", nil)
		return
	}
	rest.Success(c, "获取成功！", tag)
}

func (rest *Tag) Delete(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	db := util.GetDb(c)

	var tag model.PortalTag

	tx := db.Where("id", rewrite.Id).First(&tag)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	if tx.RowsAffected == 0 {
		rest.Error(c, "内容不存在！", nil)
		return
	}

	tx = db.Where("id", rewrite.Id).Delete(&tag)
	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, "删除成功！", nil)

}
