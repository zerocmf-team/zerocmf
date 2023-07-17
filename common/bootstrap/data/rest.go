/**
** @创建时间: 2022/3/13 18:03
** @作者　　: return
** @描述　　:
 */

package data

type Rest struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	StatusCode *int        `json:"-"`
}

type H map[string]interface{}

func (r *Rest) Success(msg string, data interface{}) (resp *Rest) {
	r.Code = 1
	r.Msg = msg
	r.Data = data
	resp = r
	return
}

func (r *Rest) Error(msg string, data interface{}) (resp *Rest) {
	r.Code = 0
	r.Msg = msg
	r.Data = data
	resp = r
	return
}
