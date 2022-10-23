/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"zerocmf/common/bootstrap/database"
)

func Migrate(tenantId string) {
	curDb := database.Conf().ManualDb(tenantId)
	new(Option).AutoMigrate(curDb)
}
