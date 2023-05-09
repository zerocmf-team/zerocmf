/**
** @创建时间: 2023/05/06 12:14
** @作者　　: return
** @描述　　:
 */

package model

import "zerocmf/common/bootstrap/database"

func Migrate(tenantId string) {
	curDb := database.Conf().ManualDb(tenantId)
	new(Site).AutoMigrate(curDb)
	new(User).AutoMigrate(curDb)
	new(Tenant).AutoMigrate(curDb)
}
