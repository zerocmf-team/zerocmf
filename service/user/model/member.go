package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Member struct {
	Id            int     `json:"id"`
	Gender        int     `gorm:"type:tinyint(2);default:0;comment:（性别，0：保密，1：男，2：女）" json:"gender"`
	Birthday      int64   `gorm:"type:bigint(20);comment:用户生日" json:"birthday"`
	BirthdayTime  string  `gorm:"-" json:"birthday_time"`
	Score         int     `gorm:"type:bigint(20);default:0;not null;comment:积分" json:"score"`
	Coin          int     `gorm:"type:bigint(20);default:0;not null;comment:金币" json:"coin"`
	Exp           int     `gorm:"type:bigint(20);default:0;not null;comment:经验" json:"exp"`
	Balance       float64 `gorm:"type:decimal(8,2);not null;comment:余额" json:"balance"`
	UserPass      string  `gorm:"type:varchar(64);comment:登录密码" json:"-"`
	UserNickname  string  `gorm:"type:varchar(50);column:user_nickname;comment:用户昵称" json:"userNickName"`
	UserRealName  string  `gorm:"type:varchar(50);column:user_realname;comment:真实姓名" json:"userRealName"`
	Avatar        string  `gorm:"type:varchar(255);comment:用户头像" json:"avatar"`
	AvatarPrev    string  `gorm:"-" json:"avatarPrev"`
	Signature     string  `gorm:"type:varchar(100);comment:用户签名" json:"signature"`
	LastLoginIp   string  `gorm:"type:varchar(50);column:last_loginip;comment:最后一次登录ip" json:"lastLoginIp"`
	Mobile        string  `gorm:"type:varchar(20);not null;comment:用户手机号" json:"mobile"`
	LastLoginAt   int64   `gorm:"type:bigint(20);comment:最后登录时间" json:"lastLoginAt"`
	CreatedAt     int64   `gorm:"type:bigint(20);comment:创建时间" json:"createdAt"`
	UpdatedAt     int64   `gorm:"type:bigint(20);comment:更新时间" json:"updatedAt"`
	LastLoginTime string  `gorm:"-" json:"lastLoginTime"`
	CreatedTime   string  `gorm:"-" json:"createdTime"`
	UpdatedTime   string  `gorm:"-" json:"updatedTime"`
	DeleteAt      int64   `gorm:"type:bigint(20);comment:删除时间" json:"deletedAt"`
	UserStatus    int     `gorm:"type:tinyint(3);not null;default:1;comment:用户状态" json:"userStatus"`
}

func (m *Member) AutoMigrate(db *gorm.DB) {
	fmt.Println("AutoMigrate")
	db.AutoMigrate(&m)
}
