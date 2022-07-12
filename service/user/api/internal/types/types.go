// Code generated by goctl. DO NOT EDIT.
package types

import (
	bsData "zerocmf/common/bootstrap/data"
	"github.com/jinzhu/copier"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ListReq struct {
	UserLogin    string `json:"user_login,optional"`
	UserNickname string `json:"userNickname,optional"`
	UserEmail    string `json:"userEmail,optional"`
}

type OneReq struct {
	Id string `path:"id,optional"`
}

type AdminStoreReq struct {
	UserLogin    string   `json:"user_login,optional"`
	UserPass     string   `json:"user_pass,optional"`
	UserEmail    string   `json:"user_email,optional"`
	Mobile       string   `json:"mobile,optional"`
	UserRealname string   `json:"user_realname,optional"`
	RoleIds      []string `json:"role_ids,optional"`
}

type AdminSaveReq struct {
	Id           string   `path:"id,optional"`
	UserLogin    string   `json:"user_login,optional"`
	UserPass     string   `json:"user_pass,optional"`
	UserEmail    string   `json:"user_email,optional"`
	Mobile       string   `json:"mobile,optional"`
	UserRealname string   `json:"user_realname,optional"`
	RoleIds      []string `json:"role_ids,optional"`
}

type RoleGet struct {
	Status string `json:"status,optional"`
	Name   string `json:"name,optional"`
}

type RoleDelete struct {
	Id  string   `path:"id,optional"`
	Ids []string `form:"ids,optional"`
}

type AccessStore struct {
	Name       string   `json:"name,optional"`
	Remark     string   `json:"remark,optional"`
	RoleAccess []string `json:"role_access,optional"`
}

type AccessEdit struct {
	Id         string   `path:"id,optional"`
	Name       string   `json:"name"`
	Remark     string   `json:"remark"`
	RoleAccess []string `json:"role_access"`
}

type AppSaveReq struct {
	Gender       int    `json:"gender,optional"`
	BirthdayTime string `json:"birthday_time,optional"`
	UserPass     string `json:"user_pass,optional"`
	UserNickname string `json:"user_nickname,optional"`
	UserRealName string `json:"user_realname,optional"`
	UserEmail    string `json:"user_email,optional"`
	UserUrl      string `json:"user_url,optional"`
	Avatar       string `json:"avatar,optional"`
	Signature    string `json:"signature,optional"`
	Mobile       string `json:"mobile,optional"`
}

type TokenReq struct {
	Usermame string `json:"username,optional"`
	Password string `json:"password,optional"`
}

type RefreshReq struct {
	RefreshToken string `json:"refreshToken,optional"`
}

type ValidationReq struct {
	TenantId string `form:"tenant_id,optional"`
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
