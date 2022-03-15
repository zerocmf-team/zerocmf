/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	new(AdminMenu).AutoMigrate(db)
	new(Option).AutoMigrate(db)
	new(Assets).AutoMigrate(db)
	//new(Comment).AutoMigrate(db)
}
