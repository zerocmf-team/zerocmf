/**
** @创建时间: 2022/3/13 22:04
** @作者　　: return
** @描述　　:
 */

package data

import (
	"errors"
	"net/http"
	"strconv"
)

type Paginate struct {
	Data     interface{} `json:"data"`
	Current  int         `json:"current"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
}

func (paginate *Paginate) Default(r *http.Request) (current int, pageSize int, err error) {

	r.ParseForm()

	qCurrent := r.Form.Get("current")
	if qCurrent == "" {
		qCurrent = "1"
	}
	qPageSize := r.Form.Get("pageSize")
	if qPageSize == "" {
		qPageSize = "10"
	}

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