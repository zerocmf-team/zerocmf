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
