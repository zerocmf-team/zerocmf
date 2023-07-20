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

	new(Option).AutoMigrate(curDb)
	new(PortalPost).AutoMigrate(curDb)
	new(PortalCategories).AutoMigrate(curDb)
	new(PortalTag).AutoMigrate(curDb)
	new(Theme).AutoMigrate(curDb)
	new(Route).AutoMigrate(curDb)
	new(Nav).AutoMigrate(curDb)

	// 评论数据库迁移
	new(Comment).AutoMigrate(curDb)

	new(App).AutoMigrate(curDb)
	new(AppPage).AutoMigrate(curDb)

	new(Form).AutoMigrate(curDb)
	new(FormItem).AutoMigrate(curDb)

}
