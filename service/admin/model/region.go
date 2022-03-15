/**
** @创建时间: 2021/11/26 10:46
** @作者　　: return
** @描述　　:
 */

package model

import (
	"github.com/gincmf/bootstrap/db"
)

type Region struct {
	AreaId   int    `gorm:"type:mediumint;primaryKey;not null;comment:行政编码" json:"area_id"`
	AreaName string `gorm:"type:varchar(20);not null;comment:地区名称" json:"area_name"`
	AreaType int    `gorm:"type:tinyint(3);not null;comment:地区类型" json:"area_type"`
	ParentId int    `gorm:"type:mediumint;not null;default:0;comment:父级行政编码" json:"parent_id"`
}

func (_ Region) AutoMigrate() {
	db := db.Db()
	db.AutoMigrate(&Region{})
}
