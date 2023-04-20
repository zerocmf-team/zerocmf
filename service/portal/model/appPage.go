package model

import (
	"errors"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
)

type AppPage struct {
	Id             int     `json:"id"`
	AppId          int     `gorm:"type:int(11);comment:渠道id;not null" json:"appId"`
	IsHome         int     `gorm:"type:tinyint(3);comment:是否是首页;default:0;not null" json:"isHome"`
	IsPublic       int     `gorm:"type:tinyint(3);comment:是否公共部分;default:0;not null" json:"isPublic"`
	Name           string  `gorm:"type:varchar(20);comment:页面名称;not null" json:"name"`
	Description    string  `gorm:"type:varchar(255);comment:页面描述;not null" json:"description"`
	Type           string  `gorm:"type:varchar(20);comment:模板类型(page：页面，list：列表);default:page;not null" json:"type"`
	SeoTitle       string  `gorm:"type:varchar(100);comment:三要素标题;not null" json:"title"`
	SeoKeywords    string  `gorm:"type:varchar(255);comment:三要素关键字;not null" json:"seoKeywords"`
	SeoDescription string  `gorm:"type:varchar(255);comment:三要素描述;not null" json:"seoDescription"`
	Schema         string  `gorm:"type:longtext;comment:页面低代码文件" json:"schema"`
	ListOrder      float64 `gorm:"type:float;comment:排序;default:10000" json:"listOrder"`
	UserId         int     `gorm:"type:int(11);comment:用户id;NOT NULL" json:"user_id"`
	CreateAt       int64   `gorm:"type:bigint(20)" json:"createAt"`
	UpdateAt       int64   `gorm:"type:bigint(20)" json:"updateAt"`
	Status         int     `gorm:"type:tinyint(3);default:1" json:"status"`
	DeleteAt       int64   `gorm:"type:int(20);comment:删除时间;default:0" json:"deleteAt"`
}

func (model *AppPage) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model *AppPage) Index(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (data.Paginate, error) {
	// 获取默认的系统分页
	var total int64 = 0
	var appPage []AppPage
	db.Where(query, queryArgs...).Find(&appPage).Count(&total)
	tx := db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&appPage)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return data.Paginate{}, tx.Error
		}
	}
	paginate := data.Paginate{Data: appPage, Current: current, PageSize: pageSize, Total: total}
	if len(appPage) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}

func (model *AppPage) List(db *gorm.DB, query string, queryArgs []interface{}) (data []AppPage, err error) {

	tx := db.Where(query, queryArgs...).Find(&data)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	return
}

func (model *AppPage) Show(db *gorm.DB, query string, queryArgs []interface{}) error {

	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}
