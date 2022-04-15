/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import "gincmf/common/bootstrap/db"

func Migrate(tenantId string) {
	curDb := db.Conf().ManualDb(tenantId)
	new(User).AutoMigrate(curDb)
	new(Role).AutoMigrate(curDb)
}
