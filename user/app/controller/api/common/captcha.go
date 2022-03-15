/**
** @创建时间: 2021/12/19 09:26
** @作者　　: return
** @描述　　:
 */

package common

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
)

type Captcha struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取验证码
 * @Date 2021/12/19 8:48:55
 * @Param
 * @return
 **/

func (rest *Captcha) New(c *gin.Context) {
	captchaId := captcha.NewLen(4)
	rest.Success(c, "获取成功", gin.H{
		"captchaId": captchaId,
		"img":       "/api/v1/app/user/captcha/" + captchaId + ".png",
	})
}

func (rest *Captcha) Captcha(c *gin.Context) {
	handler := captcha.Server(captcha.StdWidth, captcha.StdHeight)
	handler.ServeHTTP(c.Writer, c.Request)
}
