/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

func Migrate(curDb *gorm.DB) {
	new(AdminMenu).AutoMigrate(curDb)
	new(Option).AutoMigrate(curDb)
	new(Assets).AutoMigrate(curDb)
	new(model.Route).AutoMigrate(curDb)
	//new(Comment).AutoMigrate(db)
}
