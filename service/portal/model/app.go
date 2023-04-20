package model

import (
	"errors"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/data"
)

type App struct {
	Id          int     `json:"id"`
	Name        string  `gorm:"type:varchar(40);comment:应用名称;not null" json:"name"`
	Version     string  `gorm:"type:varchar(10);comment:应用版本;not null" json:"version"`
	Thumbnail   string  `gorm:"type:varchar(255);comment:主题缩略图;not null" json:"thumbnail"`
	Description string  `gorm:"type:varchar(255);comment:主题描述;not null" json:"description"`
	UserId      int     `gorm:"type:int(11);comment:用户id;NOT NULL" json:"user_id"`
	CreateAt    int64   `gorm:"type:bigint(20);comment:创建时间;default:0" json:"createAt"`
	UpdateAt    int64   `gorm:"type:bigint(20);comment:更新时间;default:0" json:"updateAt"`
	ListOrder   float64 `gorm:"type:float;comment:排序;default:10000" json:"listOrder"`
	DeleteAt    int64   `gorm:"type:int(20);comment:删除时间;default:0" json:"deleteAt"`
}

func (model *App) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
}

func (model *App) Index(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (data.Paginate, error) {
	// 获取默认的系统分页
	var total int64 = 0
	var app []App
	db.Where(query, queryArgs...).Find(&app).Count(&total)
	tx := db.Where(query, queryArgs...).Limit(pageSize).Offset((current - 1) * pageSize).Find(&app)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return data.Paginate{}, tx.Error
		}
	}
	paginate := data.Paginate{Data: app, Current: current, PageSize: pageSize, Total: total}
	if len(app) == 0 {
		paginate.Data = make([]string, 0)
	}
	return paginate, nil
}
