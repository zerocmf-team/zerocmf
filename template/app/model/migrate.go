/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/db"
)

func init() {
	conf := config.Config()
	// 单体应用，直接初始化数据
	if conf.App.Type == "single" {
		Migrate("")
	}
}

func Migrate(tenantId string) {
	db := db.ManualDb(tenantId)
	new(AdminMenu).AutoMigrate(db)
}
