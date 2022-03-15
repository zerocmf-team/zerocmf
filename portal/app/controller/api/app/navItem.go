/**
** @创建时间: 2022/1/23 16:49
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
)

type NavItem struct {
	controller.Rest
}

func (rest *NavItem) GetNavItems(c *gin.Context) {
	navId := c.Param("id")
	itemQuery := "nav_id = ? AND status = 1"
	itemQueryArgs := []interface{}{navId}
	db := util.GetDb(c)
	navItems, err := new(model.NavItem).GetWithChild(db, itemQuery, itemQueryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "操作成功！",navItems)
}