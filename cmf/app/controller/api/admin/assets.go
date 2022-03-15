/**
** @创建时间: 2021/12/6 14:09
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"gincmf/app/service"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"strings"
)

type Assets struct {
	controller.Rest
}

func (rest *Assets) Get(c *gin.Context) {
	db := util.GetDb(c)
	userId, _ := c.Get("userId")
	query := []string{"user_id = ? AND status = ?"}
	queryArgs := []interface{}{userId, "1"}
	paramType := c.DefaultQuery("type", "0")
	query = append(query, "type = ?")
	queryArgs = append(queryArgs, paramType)
	queryStr := strings.Join(query, " AND ")
	result, err := new(service.Assets).Get(c, db, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, "系统出错", err.Error())
		return
	}
	rest.Success(c, "获取成功！", result)
}

func (rest *Assets) Store(c *gin.Context) {
	db := util.GetDb(c)
	result ,err := new(service.Assets).Store(c,db)
	if err != nil {
		rest.Error(c, "系统出错", err.Error())
		return
	}
	rest.Success(c, "获取成功！", result)
}
