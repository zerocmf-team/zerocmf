/**
** @创建时间: 2022/4/4 18:59
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	bspModel "zerocmf/common/bootstrap/model"
)

type Route struct {
	bspModel.Route
}

func (model *Route) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&bspModel.Route{})
}
