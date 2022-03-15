/**
** @创建时间: 2021/11/24 22:44
** @作者　　: return
** @描述　　:
 */

package grpc

import "github.com/gincmf/bootstrap/consul"


func init() {
	registerConsul()
}

func registerConsul()  {
	consul.Register()
}