/**
** @创建时间: 2021/11/24 19:08
** @作者　　: return
** @描述　　:
 */

package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/model"
	"zerocmf/common/bootstrap/util"
)

type User struct {
	Id            int     `json:"id"`
	UserType      int     `gorm:"type:tinyint(3);not null;default:0;comment:（管理员类型，0：普通用户，1：系统管理员）" json:"user_type"`
	Gender        int     `gorm:"type:tinyint(2);default:0;comment:（性别，0：保密，1：男，2：女）" json:"gender"`
	Birthday      int64   `gorm:"type:bigint(20);comment:用户生日" json:"birthday"`
	BirthdayTime  string  `gorm:"-" json:"birthday_time"`
	Score         int     `gorm:"type:bigint(20);default:0;not null;comment:积分" json:"score"`
	Coin          int     `gorm:"type:bigint(20);default:0;not null;comment:金币" json:"coin"`
	Exp           int     `gorm:"type:bigint(20);default:0;not null;comment:经验" json:"exp"`
	Balance       float64 `gorm:"type:decimal(8,2);not null;comment:余额" json:"balance"`
	UserLogin     string  `gorm:"type:varchar(60);comment:登录账号" json:"user_login"`
	UserPass      string  `gorm:"type:varchar(64);comment:登录密码" json:"-"`
	UserNickname  string  `gorm:"type:varchar(50);column:user_nickname;comment:用户昵称" json:"user_nickname"`
	UserRealName  string  `gorm:"type:varchar(50);column:user_realname;comment:真实姓名" json:"user_realname"`
	UserEmail     string  `gorm:"type:varchar(100);comment:用户邮箱" json:"user_email"`
	UserUrl       string  `gorm:"type:varchar(100);comment:用户主页网站" json:"user_url"`
	Avatar        string  `gorm:"type:varchar(255);comment:用户头像" json:"avatar"`
	AvatarPrev    string  `gorm:"-" json:"avatar_prev"`
	Signature     string  `gorm:"type:varchar(100);comment:用户签名" json:"signature"`
	LastLoginIp   string  `gorm:"type:varchar(50);column:last_loginip;comment:最后一次登录ip" json:"last_loginip"`
	Mobile        string  `gorm:"type:varchar(20);not null;comment:用户手机号" json:"mobile"`
	LastLoginAt   int64   `gorm:"type:bigint(20);comment:最后登录时间" json:"last_login_at"`
	CreateAt      int64   `gorm:"type:bigint(20);comment:创建时间" json:"create_at"`
	UpdateAt      int64   `gorm:"type:bigint(20);comment:更新时间" json:"update_at"`
	LastLoginTime string  `gorm:"-" json:"last_login_time"`
	CreateTime    string  `gorm:"-" json:"create_time"`
	UpdateTime    string  `gorm:"-" json:"update_time"`
	DeleteAt      int64   `gorm:"type:bigint(20);comment:删除时间" json:"delete_at"`
	UserStatus    int     `gorm:"type:tinyint(3);not null;default:1;comment:用户状态" json:"user_status"`
	model.Time
}

type ThirdPart struct {
	Id         int    `json:"id"`
	Type       string `gorm:"type:varchar(10);not null" json:"type"`
	UserId     int    `gorm:"type:int(11);not null" json:"user_id"`
	AppId      string `gorm:"type:varchar(64);not null" json:"app_id"`
	OpenId     string `gorm:"type:varchar(20);not null" json:"open_id"`
	SessionKey string `gorm:"-" json:"session_key"`
	Status     int    `gorm:"type:tinyint(3);not null;default:1;comment:状态;1:正常;0:禁用" json:"status"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {

	if u.Avatar != "" {

		// 调用资源rpc，获取资源服务器地址

		prevPath := util.FileUrl(u.Avatar)
		if err != nil {
			fmt.Println("err", err)
		}
		u.AvatarPrev = prevPath
	}

	birthTime := time.Unix(u.Birthday, 0).Format("2006-01-02 15:04:05")
	u.BirthdayTime = birthTime
	return
}

func (u *User) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})

	tx := db.Where("user_login = ?", "admin").First(&User{}) // 查询

	if tx.RowsAffected == 0 {
		//新增一条管理员数据
		db.Create(&User{UserType: 1, UserLogin: "admin", UserPass: util.GetMd5("123456"), LastLoginAt: time.Now().Unix(), CreateAt: time.Now().Unix()})
	}
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/12/20 12:49:5
 * @Param 获取分页数据
 * @return
 **/

func (u *User) Paginate(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (result data.Paginate, err error) {
	var user []User
	var total int64 = 0
	tx := db.Where(query, queryArgs...).Find(&user).Count(&total)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	tx = db.Where(query, queryArgs...).Debug().Limit(pageSize).Offset((current - 1) * pageSize).Find(&user)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	for k, v := range user {
		if v.LastLoginAt > 0 {
			user[k].LastLoginTime = time.Unix(v.LastLoginAt, 0).Format(data.TimeLayout)
		}
		if v.CreateAt > 0 {
			user[k].CreateTime = time.Unix(v.CreateAt, 0).Format(data.TimeLayout)
		}
	}
	result = data.Paginate{Data: user, Current: current, PageSize: pageSize, Total: total}
	if len(user) == 0 {
		result.Data = make([]string, 0)
	}
	return result, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看
 * @Date 2021/12/18 13:28:39
 * @Param
 * @return
 **/

func (u *User) Show(db *gorm.DB, query string, queryArgs []interface{}) error {
	tx := db.Where(query, queryArgs...).First(&u)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 注册新用户
 * @Date 2021/12/18 12:50:17
 * @Param
 * @return
 **/

func (u *User) Register(db *gorm.DB) error {

	user := User{}
	err := user.Show(db, "mobile = ?", []interface{}{u.Mobile})
	if err != nil {
		return err
	}

	if user.Id > 0 {
		return errors.New("该用户已经注册，请直接登录")
	}

	tx := db.Create(&u)
	if tx.Error != nil {
		return tx.Error
	}

	return nil

}
