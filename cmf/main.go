/**
** @创建时间: 2021/11/23 09:26
** @作者　　: return
** @描述　　:
 */

package main

import (
	_ "gincmf/app/grpc" // 注册rpc
	_ "gincmf/app/model" // 数据库迁移
	_ "gincmf/routes"    // 路由注册
)

func main() {

}