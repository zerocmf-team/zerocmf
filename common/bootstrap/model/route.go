/**
** @创建时间: 2020/12/26 3:00 下午
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"zerocmf/common/bootstrap/util"
	"gorm.io/gorm"
)

type Route struct {
	Id           int     `json:"id"`
	ListOrder    float64 `gorm:"type:float;comment:排序;default:10000" json:"list_order"`
	Status       int     `gorm:"type:tinyint(3);comment:状态;default:1;;not null" json:"status"`
	Type         int     `gorm:"type:tinyint(4);comment:URL规则类型;1:用户自定义;2:别名添加';default:1;;not null" json:"type"`
	FullUrl      string  `gorm:"type:varchar(255);comment:完整url;not null" json:"full_url"`
	Url          string  `gorm:"type:varchar(255);comment:实际显示的url;not null" json:"url"`
}

type RouteResult struct {
	Route
	Template string `json:"template"`
}

func (model *Route) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model *Route) Show(db *gorm.DB, query string, queryArgs []interface{}) error {
	tx := db.Where(query, queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}

func (model *Route) List(db *gorm.DB, query string, queryArgs []interface{}) ([]Route, error) {
	var route []Route
	tx := db.Where(query, queryArgs...).Find(&route)
	if util.IsDbErr(tx) != nil {
		return route, nil
	}
	return route, nil
}

func (model *Route) Set(db *gorm.DB) error {
	route := Route{
		Type:         model.Type,
		FullUrl:      model.FullUrl,
		Url:          model.Url,
	}
	tx := db.Where("full_url", route.FullUrl).First(&route)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return tx.Error
	}
	if tx.RowsAffected == 0 {
		db.Create(&route)
	} else {
		route.Type = model.Type
		route.FullUrl = model.FullUrl
		route.Url = model.Url
		tx := db.Save(&route)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}
