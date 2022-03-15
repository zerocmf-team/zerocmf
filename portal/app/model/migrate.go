/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"github.com/gincmf/bootstrap/db"
)

func init() {
	Migrate("")
}

func Migrate(tenantId string) {
	db := db.ManualDb(tenantId)
	new(AdminMenu).AutoMigrate(db)
	new(PortalPost).AutoMigrate(db)
	new(PortalCategory).AutoMigrate(db)
	new(PortalTag).AutoMigrate(db)
	new(Theme).AutoMigrate(db)
	new(Nav).AutoMigrate(db)
	new(Option).AutoMigrate(db)
}
