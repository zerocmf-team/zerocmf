/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import "gincmf/common/bootstrap/db"

func Migrate(tenantId string) {
	curDb := db.Conf().ManualDb(tenantId)
	new(Option).AutoMigrate(curDb)
	new(PortalPost).AutoMigrate(curDb)
	new(PortalCategory).AutoMigrate(curDb)
	new(PortalTag).AutoMigrate(curDb)
	new(Theme).AutoMigrate(curDb)
	new(Route).AutoMigrate(curDb)
	new(Nav).AutoMigrate(curDb)
}
