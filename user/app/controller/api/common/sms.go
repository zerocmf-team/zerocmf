/**
** @创建时间: 2021/12/19 12:01
** @作者　　: return
** @描述　　:
 */

package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/sms"
)

type Sms struct {
	controller.Rest
}

var DaYu = sms.DaYu{
	RegionId:        "cn-hangzhou",
	AccessKeyId:     "您的阿里云AccessKeyId",
	AccessKeySecret: "您的阿里云AccessKeySecret",
}

func (rest *Sms) Send(c *gin.Context) {
	mobile := c.Query("mobile")
	scene := c.Query("scene")

	if mobile == "" {
		rest.Error(c, "手机号不能为空！", nil)
		return
	}

	if scene == "" {
		rest.Error(c, "场景不能为空！", nil)
		return
	}

	DaYu.Scheme = "http"
	DaYu.SignName = "码上云"
	DaYu.TemplateCode = "SMS_203678233"
	code, err := DaYu.Send(mobile, scene)
	if err != nil {
		rest.Error(c, "发送失败！", nil)
		return
	}
	rest.Success(c, "发送成功！", code)
}

func (rest *Sms) Verify(c *gin.Context) {
	code := c.Query("code")
	err := DaYu.Verify(code, "user/register")
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "校验成功！", nil)
}
