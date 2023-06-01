/**
** @创建时间: 2023/05/06 12:14
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	new(Site).AutoMigrate(db)
	new(User).AutoMigrate(db)
	new(Tenant).AutoMigrate(db)
}
