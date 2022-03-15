/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"github.com/gincmf/bootstrap/db"
	"github.com/gincmf/bootstrap/model"
)

func init() {
	// 主程序应用，直接初始化数据
	Migrate("")
}

func Migrate(tenantId string) {
	db := db.ManualDb(tenantId)
	new(model.Route).AutoMigrate(db)
	new(AdminMenu).AutoMigrate(db)
	new(Option).AutoMigrate(db)
	new(Assets).AutoMigrate(db)
	new(Comment).AutoMigrate(db)
}
