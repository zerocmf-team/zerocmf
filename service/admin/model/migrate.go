/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gincmf/common/bootstrap/db"
	"gincmf/common/bootstrap/model"
)

func Migrate(tenantId string,api bool) {
	curDb := db.Conf().ManualDb(tenantId)
	if api {
		new(AdminMenu).AutoMigrate(curDb)
	}
	new(Option).AutoMigrate(curDb)
	new(Assets).AutoMigrate(curDb)
	new(model.Route).AutoMigrate(curDb)
	//new(Comment).AutoMigrate(db)
}
