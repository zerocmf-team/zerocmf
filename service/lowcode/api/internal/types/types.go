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

type FormGetReq struct {
}

type FormShowReq struct {
	FormId string `path:"formId,optional"`
}

type FormSaveReq struct {
	Id          string  `json:"id,optional"`
	ParentId    string  `json:"parentId,optional"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon,optional"`
	MenuType    int     `json:"menuType,optional"`
	HideInMenu  int     `json:"hideInMenu,optional"`
	Description string  `json:"description,optional"`
	Schema      string  `json:"schema,optional"`
	ListOrder   float64 `json:"listOrder,optional"`
	Status      int     `json:"status,optional"`
}

type FormDatasReq struct {
	FormId   string `form:"formId,optional"`
	Current  *int   `form:"current,optional"`
	PageSize *int   `form:"pageSize,optional"`
}

type FormDataShowReq struct {
	Id string `path:"id"`
}

type FormDataSaveReq struct {
	Id           string `path:"id,optional"`
	FormId       string `json:"formId"`
	FormDataJson string `json:"formDataJson"`
}

type RegionGetReq struct {
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