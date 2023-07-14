/**
** @创建时间: 2022/3/13 22:04
** @作者　　: return
** @描述　　:
 */

package data

import (
	"net/http"
	"strconv"
)

type Paginate struct {
	Data     interface{} `bson:"data" json:"data"`
	Current  int         `bson:"current" json:"current"`
	PageSize int         `bson:"pageSize" json:"pageSize"`
	Total    int64       `bson:"total" json:"total"`
}

type paginate struct {
	Request *http.Request `json:"-"`
}

func NewPaginate(req *http.Request) (p *paginate) {
	p = new(paginate)
	p.Request = req
	return
}

/**
 * @Author return
 * @Description 获取请求中的当前页码和分页数
 * @Date 2022/12/1 23:21
 * @Param
 * @return
 **/

func (page *paginate) Default() (current int, pageSize int, err error) {
	r := page.Request
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
		current = 1
	}

	return current, pageSize, nil
}
