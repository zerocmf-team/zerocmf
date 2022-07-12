/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"zerocmf/common/bootstrap/database"
	"zerocmf/common/bootstrap/model"
)

func Migrate(tenantId string,api bool) {
	curDb := database.Conf().ManualDb(tenantId)
	if api {
		new(AdminMenu).AutoMigrate(curDb)
	}
	new(Option).AutoMigrate(curDb)
	new(Assets).AutoMigrate(curDb)
	new(model.Route).AutoMigrate(curDb)
	//new(Comment).AutoMigrate(db)
}
