package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

type Tenant struct {
	Id       int64  `json:"id"`
	TenantId int64  `gorm:"type:bigint(20);not null;index:idx_tenant_id" json:"tenantId"`
	Company  string `gorm:"type:varchar(100);not null;comment:公司名称" json:"company"`
	Mobile   string `gorm:"type:varchar(20);not null" json:"mobile"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Status   int    `gorm:"type:tinyint(3);comment:租户状态;default:1;not null" json:"status"`
	model.Time
}

type TenantUser struct {
	Id       int64 `json:"id"`
	TenantId int64 `gorm:"type:bigint(20);comment:租户id;not null" json:"tenantId"`
	Uid      int64 `gorm:"type:bigint(20);comment:统一站点用户id;not null" json:"uid"`
	Status   int   `gorm:"type:tinyint(3);not null;default:1;comment:状态;1:正常;0:禁用" json:"status"`
}

func (model *Tenant) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
	db.AutoMigrate(&TenantUser{})
}
