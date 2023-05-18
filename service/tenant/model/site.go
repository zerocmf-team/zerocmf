package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

type Site struct {
	Id     int64  `json:"id"`
	SiteId int64  `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
	Name   string `gorm:"type:varchar(32);comment:站点名称" json:"name"`
	Domain string `gorm:"type:varchar(100);comment:站点域名" json:"domain"`
	Desc   string `gorm:"type:varchar(255);comment:站点描述" json:"desc"`
	Dsn    string `gorm:"type:varchar(32);comment:数据库配置" json:"dsn"`
	Status int    `gorm:"type:tinyint(3);default:1;comment:文件状态" json:"status"`
	model.Time
	DeleteAt   int64  `gorm:"type:bigint(20);comment:删除实际;NOT NULL" json:"delete_at"`
	DeleteTime string `gorm:"-" json:"delete_time"`
}

type SiteUser struct {
	Id        int64   `json:"id"`
	TenantId  int64   `gorm:"type:bigint(20);comment:租户id;not null" json:"tenantId"`
	SiteId    int64   `gorm:"type:bigint(20);comment:站点id;not null" json:"siteId"`
	Uid       int64   `gorm:"type:bigint(20);comment:统一站点唯一用户id;not null" json:"uid"`
	Oid       int64   `gorm:"type:bigint(20);comment:真实站点用户id;not null" json:"oid"`
	IsOwner   int     `gorm:"type:tinyint(3);comment:是否为站点拥有者;not null" json:"isOwner"`
	ListOrder float64 `gorm:"type:float;default:10000;comment:排序（越大越靠前）" json:"listOrder" label:"排序"`
	Status    int     `gorm:"type:tinyint(3);not null;default:1;comment:状态;1:正常;0:禁用" json:"status"`
}

func (m Site) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&m)
}
