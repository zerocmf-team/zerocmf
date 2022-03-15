/**
** @创建时间: 2021/12/1 11:26
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type User struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 根据token获取当前用户信息
 * @Date 2021/12/1 11:27:15
 * @Param
 * @return
 **/

func (rest *User) CurrentUser(c *gin.Context) {
	// 获取当前用户
	userId, _ := c.Get("userId")
	db := util.GetDb(c)

	user := model.User{}
	tx := db.Where("id = ? and user_type = 1", userId).First(&user)

	if tx.Error != nil {
		err := "系统错误：" + tx.Error.Error()
		if tx.Error == gorm.ErrRecordNotFound {
			err = "该管理员账号不存在"
		}
		rest.Error(c, err, nil)
		return
	}

	rest.Success(c, "获取成功！", user)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 修改个人信息
 * @Date 2022/3/4 12:36:20
 * @Param
 * @return
 **/

func (rest *User) Save(c *gin.Context) {

	userId, _ := c.Get("userId")
	db := util.GetDb(c)
	user := model.User{}
	err := user.Show(db, "id = ?", []interface{}{userId})
	if err != nil {
		msg := "系统错误：" + err.Error()
		if err == gorm.ErrRecordNotFound {
			msg = "该管理员账号不存在"
		}
		rest.Error(c, msg, nil)
		return
	}

	var form struct {
		Gender       int    `json:"gender"`
		BirthdayTime string `json:"birthday_time"`
		UserPass     string `json:"user_pass"`
		UserNickname string `json:"user_nickname"`
		UserRealName string `json:"user_realname"`
		UserEmail    string `json:"user_email"`
		UserUrl      string `json:"user_url"`
		Avatar       string `json:"avatar"`
		Signature    string `json:"signature"`
		Mobile       string `json:"mobile"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	copier.Copy(&user, &form)

	if form.BirthdayTime != "" {
		times, _ := time.Parse("2006-01-02", form.BirthdayTime)
		birthday := times.Unix()
		user.Birthday = birthday
	}

	tx := db.Where("id = ?", userId).Save(&user)

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	user.AvatarPrev = util.FileUrl(user.Avatar)

	rest.Success(c, "操作成功！", user)
}
