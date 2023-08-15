/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/model"
)

// 管理员表

type User struct {
	Id            int64  `json:"id"`
	Uid           int64  `gorm:"type:bigint(20);comment:唯一用户id" json:"uid"`
	UserLogin     string `gorm:"type:varchar(60);comment:登录账号" json:"user_login"`
	Mobile        string `gorm:"type:varchar(20);not null" json:"mobile"`
	UserRealName  string `gorm:"type:varchar(20);not null" json:"userRealName"`
	UserPass      string `gorm:"type:varchar(64);comment:登录密码" json:"-"`
	LastLoginIp   string `gorm:"type:varchar(50);comment:最后一次登录ip" json:"lastLoginIp"`
	LastLoginAt   int64  `gorm:"type:bigint(20);comment:最后登录时间" json:"lastLoginAt"`
	LastLoginTime string `gorm:"-" json:"last_loginTime"`
	DeleteAt      int64  `gorm:"type:bigint(20);comment:删除时间" json:"deleteAt"`
	Access        int    `gorm:"-" json:"access"`
	model.Time
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

func (u *User) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SiteUser{})

	/*tx := db.Where("user_login = ?", "admin").First(&User{}) // 查询
	if tx.RowsAffected == 0 {
		//新增一条管理员数据
		db.Create(&User{Uid: 1, UserLogin: "admin", UserPass: util.GetMd5("123456"), LastLoginAt: time.Now().Unix(), Time: model.Time{
			CreateAt: time.Now().Unix(),
		}})
	}*/
}
