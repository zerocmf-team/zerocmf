/**
** @创建时间: 2021/1/3 11:23 下午
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/model"
)

type Route struct {
	controller.Rest
}

func (rest *Route) List(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		rest.Error(c,err.Error(),nil)
		return
	}
	db := util.GetDb(c)
	data,err := new(model.Route).List(db,"",nil)
	if err != nil {
		rest.Error(c,err.Error(),nil)
		return
	}
	rest.Success(c,"获取成功！",data)
}
