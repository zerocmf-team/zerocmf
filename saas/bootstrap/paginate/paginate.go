/**
** @创建时间: 2021/11/23 14:34
** @作者　　: return
** @描述　　:
 */

package paginate

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

func (paginate *Paginate) Default(c *gin.Context) (current int, pageSize int, err error) {

	qCurrent := c.DefaultQuery("current", "1")
	qPageSize := c.DefaultQuery("pageSize", "10")

	current, _ = strconv.Atoi(qCurrent)
	pageSize, _ = strconv.Atoi(qPageSize)

	if current <= 0 {
		return 0,0,errors.New("当前页码需大于0！")
	}

	if pageSize <= 0 {
		return 0,0,errors.New("每页数需大于0！")
	}

	return current,pageSize,nil
}