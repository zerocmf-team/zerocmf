/**
** @创建时间: 2022/3/13 22:04
** @作者　　: return
** @描述　　:
 */

package data

import (
	"net/http"
	"strconv"
	"zerocmf/common/bootstrap/util"
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

func PaginateQuery(r *http.Request) (current, pageSize int, err error) {
	queryParams := r.URL.Query()
	qCurrent := queryParams.Get("current")
	if qCurrent == "" {
		qCurrent = "1"
	}
	qPageSize := queryParams.Get("pageSize")
	if qPageSize == "" {
		qPageSize = "10"
	}

	current, err = strconv.Atoi(qCurrent)
	if err != nil {
		current = 1
	}

	pageSize, err = strconv.Atoi(qPageSize)
	if err != nil {
		pageSize = 10
	}

	if current <= 0 {
		current = 1
	}
	return
}

func PaginateQueryInt32(r *http.Request) (current, pageSize int32, err error) {
	c, size, qErr := PaginateQuery(r)
	if qErr != nil {
		return 1, 10, nil
	}

	current, err = util.SafeToInt32(c)
	if err != nil {
		return 1, 10, nil
	}
	pageSize, err = util.SafeToInt32(size)
	if err != nil {
		return 1, 10, nil
	}

	return
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

	queryParams := r.URL.Query()
	qCurrent := queryParams.Get("current")
	if qCurrent == "" {
		qCurrent = "1"
	}
	qPageSize := queryParams.Get("pageSize")

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
