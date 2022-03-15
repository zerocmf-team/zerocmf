/**
** @创建时间: 2021/12/19 12:29
** @作者　　: return
** @描述　　: 阿里大于
 */

package sms

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"time"
)

type DaYu struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	Scheme          string
	SignName        string
	TemplateCode    string
	TemplateParam   string
}

type smsCode struct {
	Mobile string `json:"mobile"`
	Scene  string `json:"scene"`
	Code   string `json:"code"`
	Expire int64  `json:"expire"`
}

var smsArr = make(map[string]*smsCode, 0)

func (dy *DaYu) Send(mobile string, scene string) (s *smsCode, err error) {

	if mobile == "" {
		return nil, errors.New("手机号不能为空！")
	}

	if scene == "" {
		return nil, errors.New("场景不能为空！")
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%04v", rnd.Int31n(10000))
	if smsArr[code] == nil {
		smsArr[code] = new(smsCode)
	}
	if smsArr[code].Expire > time.Now().Unix() {
		return nil, errors.New("获取验证码过于频繁！请先验证上一次验证码！")
	}
	//client, err := dysmsapi.NewClientWithAccessKey(dy.RegionId, dy.AccessKeyId, dy.AccessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = mobile
	request.Scheme = dy.Scheme
	request.SignName = dy.SignName
	request.TemplateCode = dy.TemplateCode
	request.TemplateParam = `{"code":"` + code + `"}`
	//_, err = client.SendSms(request)
	//if err != nil {
	//	return nil, errors.New(err.Error())
	//}
	smsArr[code] = &smsCode{
		Mobile: mobile,
		Scene:  scene,
		Code:   code,
		Expire: time.Now().Unix() + 60*2,
	}
	return smsArr[code], nil
}

func (dy *DaYu) Verify(code string, scene string) error {
	curCode := smsArr[code]
	if curCode == nil {
		goto codeExpireErr
	}
	if curCode.Expire < time.Now().Unix() {
		goto codeExpireErr
	}
	if curCode.Code != code || curCode.Scene != scene {
		return errors.New("短信验证失败！请检查验证码是否正确")
	}

	// 验证成功
	delete(smsArr, code)

	return nil
codeExpireErr:
	return errors.New("短信验证码已经失效！请重新获取")
}
