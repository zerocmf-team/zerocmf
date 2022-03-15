/**
** @创建时间: 2021/12/18 11:09
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/controller/api/common"
	"gincmf/app/model"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"regexp"
)

type User struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 用户注册处理逻辑
 * @Date 2021/12/18 11:9:21
 * @Param
 * @return
 **/

func (rest *User) Register(c *gin.Context) {

	var form struct {
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
		Mobile    string `json:"mobile"`     // 手机号
		CaptchaId string `json:"captcha_id"` // 验证码id
		Captcha   string `json:"captcha"`    //图形验证码
		Code      string `json:"code"`       // 手机号验证码
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if form.Nickname == "" {
		rest.Error(c, "昵称不能为空！", nil)
		return
	}

	if form.Password == "" {
		rest.Error(c, "密码不能为空！", nil)
		return
	}

	if form.Mobile == "" {
		rest.Error(c, "手机号不能为空！", nil)
		return
	}

	reg := regexp.MustCompile(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`)
	matched := reg.MatchString(form.Mobile)
	if matched == false {
		rest.Error(c, "手机号格式校验失败！", nil)
		return
	}

	captchaId := form.CaptchaId
	inCaptcha := form.Captcha

	if captchaId == "" {
		rest.Error(c, "图形验证码不能为空！", nil)
		return
	}

	code := form.Code
	if code == "" {
		rest.Error(c, "请先获取手机验证码", nil)
		return
	}

	verify := captcha.VerifyString(captchaId, inCaptcha)
	if !verify {
		rest.Error(c, "图形验证码输入错误", nil)
		return
	}

	scene := "user/register"
	err := common.DaYu.Verify(code, scene)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	userPass := util.GetMd5(form.Password)
	user := model.User{
		UserType:     0,
		UserLogin:    form.Nickname,
		UserNickname: form.Nickname,
		Mobile:       form.Mobile,
		UserPass:     userPass,
	}
	db := util.GetDb(c)
	err = user.Register(db)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "注册成功！", nil)
}
