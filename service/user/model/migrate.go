/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
)

func Migrate(curDb *gorm.DB) {
	new(User).AutoMigrate(curDb)
	new(Role).AutoMigrate(curDb)
	new(Department).AutoMigrate(curDb)
}
