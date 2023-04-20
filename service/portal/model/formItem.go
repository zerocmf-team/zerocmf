package model

import "gorm.io/gorm"

type FormItem struct {
	Id       int64  `json:"id"`
	FormId   int64  `json:"form_id"`
	UserId   int    `gorm:"type:int(11);comment:发表者用户id" json:"user_id"`
	Schema   string `gorm:"type:longtext;comment:页面低代码文件" json:"schema"`
	CreateAt int64  `gorm:"type:bigint(20)" json:"createAt"`
	UpdateAt int64  `gorm:"type:bigint(20)" json:"updateAt"`
	Status   int    `gorm:"type:tinyint(3);default:1" json:"status"`
	DeleteAt int64  `gorm:"type:int(20);comment:删除时间;default:0" json:"deleteAt"`
}

func (model *FormItem) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}
