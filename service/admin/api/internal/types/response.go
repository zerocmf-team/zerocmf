/**
** @创建时间: 2022/3/16 20:26
** @作者　　: return
** @描述　　:
 */

package types

import (
	bsData "gincmf/common/bootstrap/data"
	"github.com/jinzhu/copier"
)

func (r *Response) Success(msg string, data interface{}) {
	success := new(bsData.Rest).Success(msg,data)
	copier.Copy(&r,&success)
	return
}

func (r *Response) Error(msg string, data interface{}) {
	err := new(bsData.Rest).Error(msg,data)
	copier.Copy(&r,&err)
	return
}