package model

import (
	"errors"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
)

type Form struct {
	Id             int64   `json:"id"`
	Name           string  `gorm:"type:varchar(20);comment:表单名称;not null" json:"name"`
	Description    string  `gorm:"type:varchar(255);comment:表单描述;not null" json:"description"`
	SeoTitle       string  `gorm:"type:varchar(100);comment:三要素标题;not null" json:"title"`
	SeoKeywords    string  `gorm:"type:varchar(255);comment:三要素关键字;not null" json:"seoKeywords"`
	SeoDescription string  `gorm:"type:varchar(255);comment:三要素描述;not null" json:"seoDescription"`
	Schema         string  `gorm:"type:longtext;comment:页面低代码文件" json:"schema"`
	Columns        string  `gorm:"type:longtext;comment:表头标识文件" json:"columns"`
	ListOrder      float64 `gorm:"type:float;comment:排序;default:10000" json:"listOrder"`
	UserId         int     `gorm:"type:int(11);comment:用户id;NOT NULL" json:"user_id"`
	CreateAt       int64   `gorm:"type:bigint(20)" json:"createAt"`
	UpdateAt       int64   `gorm:"type:bigint(20)" json:"updateAt"`
	Status         int     `gorm:"type:tinyint(3);default:1" json:"status"`
	DeleteAt       int64   `gorm:"type:int(20);comment:删除时间;default:0" json:"deleteAt"`
}

func (model *Form) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model *Form) List(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}, paginate bool) (result interface{}, err error) {
	var form []Form
	if paginate == true {
		// 获取默认的系统分页
		var total int64 = 0
		tx := db.Where(query, queryArgs...).Find(&form).Count(&total)
		if tx.Error != nil {
			err = tx.Error
			return
		}
		tx = db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&form)
		if tx.Error != nil {
			if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
				err = tx.Error
				return
			}
		}
		pageData := data.Paginate{Data: form, Current: current, PageSize: pageSize, Total: total}
		if len(form) == 0 {
			pageData.Data = make([]string, 0)
		}
		result = pageData
	} else {
		tx := db.Where(query, queryArgs...).Find(&form)
		if util.IsDbErr(tx) != nil {
			err = tx.Error
			return
		}
		result = form
	}
	return
}

func (model *Form) Show(db *gorm.DB, query string, queryArgs []interface{}) error {
	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}
