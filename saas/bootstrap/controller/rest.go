/**
** @创建时间: 2021/11/24 18:42
** @作者　　: return
** @描述　　:
 */

package controller

import (
	"gincmf/bootstrap/paginate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rest struct {}

func (r Rest) Success(c *gin.Context, msg string, data interface{}) {
	var result paginate.ReturnData
	result = paginate.ReturnData{Code: 1, Msg: msg, Data: data}
	c.JSON(http.StatusOK, result)
}

func (r Rest) Error(c *gin.Context, msg string, data interface{}) {
	var result paginate.ReturnData
	result = paginate.ReturnData{Code: 0, Msg: msg, Data: data}
	c.JSON(http.StatusOK, result)
}
