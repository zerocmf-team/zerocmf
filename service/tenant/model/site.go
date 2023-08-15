package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

type Site struct {
	Id     int64  `json:"id"`
	SiteId int64  `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
	Name   string `gorm:"type:varchar(32);comment:站点名称" json:"name"`
	//Domain string `gorm:"type:varchar(100);comment:站点域名" json:"domain"`
	Desc   string `gorm:"type:varchar(255);comment:站点描述" json:"desc"`
	Dsn    string `gorm:"type:varchar(32);comment:数据库配置" json:"dsn"`
	Status int    `gorm:"type:tinyint(3);default:1;comment:文件状态" json:"status"`
	model.Time
	DeleteAt   int64  `gorm:"type:bigint(20);comment:删除实际;NOT NULL" json:"deleteAt"`
	DeleteTime string `gorm:"-" json:"deleteTime"`
}

// 站点授权

type SiteMpAuth struct {
	Id              int64  `json:"id"`
	SiteId          int64  `gorm:"type:bigint(20);comment;站点唯一编号" json:"siteId"`
	Type            string `gorm:"type:varchar(20);comment:授权商户小程序类型;not null" json:"type"`
	AuthAppId       string `gorm:"type:varchar(20);comment:授权商户的appId;not null" json:"authAppId"`
	AppAuthToken    string `gorm:"type:varchar(255);comment:应用授权令牌;not null" json:"appAuthToken"`
	AppRefreshToken string `gorm:"type:varchar(100);comment:刷新令牌;not null" json:"appRefreshToken"`
	ExpiresIn       string `gorm:"type:varchar(16);comment:应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒;not null" json:"expiresIn"`
	ReExpiresIn     string `gorm:"type:varchar(16);comment:刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒;not null" json:"reExpiresIn"`
	CreatedAt       int64  `gorm:"type:int(10);comment:创建时间;default:0" json:"createdAt"`
	UpdatedAt       int64  `gorm:"type:int(10);comment:更新时间;default:0" json:"updatedAt"`
}

func (m Site) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&m)
	db.AutoMigrate(&SiteMpAuth{})
}
