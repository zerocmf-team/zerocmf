/**
** @创建时间: 2022/3/16 12:33
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Time struct {
	CreateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	UpdateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"update_at"`
	CreateTime string `gorm:"-" json:"create_time"`
	UpdateTime string `gorm:"-" json:"update_time"`
}

func (model *Time) AfterFind(tx *gorm.DB) (err error) {
	if model.CreateAt > 0 {
		createTime := time.Unix(model.CreateAt, 0).Format("2006-01-02 15:04:05")
		model.CreateTime = createTime
	}

	if model.UpdateAt > 0 {
		updateTime := time.Unix(model.UpdateAt, 0).Format("2006-01-02 15:04:05")
		model.UpdateTime = updateTime
	}
	return
}
