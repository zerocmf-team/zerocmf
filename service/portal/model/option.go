/**
** @创建时间: 2022/2/2 21:59
** @作者　　: return
** @描述　　:
 */

package model

import "gorm.io/gorm"

type Option struct {
	Id          int    `json:"id"`
	AutoLoad    int    `gorm:"type:tinyint(3);default:1;not null" json:"autoload"`
	OptionName  string `gorm:"type:varchar(64);not null" json:"option_name"`
	OptionValue string `gorm:"type:longtext" json:"option_value"`
}

func (_ *Option) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Option{})
}
