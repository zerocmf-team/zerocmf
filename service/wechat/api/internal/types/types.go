// Code generated by goctl. DO NOT EDIT.
package types

import (
	"github.com/jinzhu/copier"
	bsData "zerocmf/common/bootstrap/data"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Code2SessionReq struct {
	JsCode string `json:"js_code,optional"`
}

type CheckSignatureReq struct {
	Signature string `form:"signature,optional"`
	Timestamp string `form:"timestamp,optional"`
	Nonce     string `form:"nonce,optional"`
	Echostr   string `form:"echostr,optional"`
}

type CheckQrcodeScanReq struct {
	SceneId string `form:"scene_id,optional"`
}

func (r *Response) Success(msg string, data interface{}) {
	success := new(bsData.Rest).Success(msg, data)
	copier.Copy(&r, &success)
	return
}

func (r *Response) Error(msg string, data interface{}) {
	err := new(bsData.Rest).Error(msg, data)
	copier.Copy(&r, &err)
	return
}
