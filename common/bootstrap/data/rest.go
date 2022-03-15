/**
** @创建时间: 2022/3/13 18:03
** @作者　　: return
** @描述　　:
 */

package data

type Rest struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Rest) Success(msg string, data interface{}) (res Rest) {
	res = Rest{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
	return
}

func (r Rest) Error(msg string, data interface{}) (res Rest) {
	res = Rest{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	return
}
