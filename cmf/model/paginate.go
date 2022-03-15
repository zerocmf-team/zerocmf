/**
** @创建时间: 2020/10/30 10:10 下午
** @作者　　: return
** @描述　　:
 */
package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Paginate struct {
	Data     interface{} `json:"data"`
	Current  int      `json:"current"`
	PageSize int      `json:"page_size"`
	Total    int64       `json:"total"`
}

type ReturnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (model *Paginate) Default(c *gin.Context) (intCurrent int, intPageSize int, err error) {

	current := c.DefaultQuery("current", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	intCurrent, _ = strconv.Atoi(current)
	intPageSize, _ = strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		return 0,0,errors.New("当前页码需大于0！")
	}

	if intPageSize <= 0 {
		return 0,0,errors.New("每页数需大于0！")
	}

	return intCurrent,intPageSize,nil
}
