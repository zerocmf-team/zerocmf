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

type InitReq struct {
}

type FormGetReq struct {
}

type FormShowReq struct {
	FormId string `path:"formId,optional"`
}

type FormSaveReq struct {
	Id          string   `json:"id,optional"`
	ParentId    *string  `json:"parentId,optional"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon,optional"`
	MenuType    *int     `json:"menuType,optional"`
	HideInMenu  *int     `json:"hideInMenu,optional"`
	Description string   `json:"description,optional"`
	Schema      string   `json:"schema,optional"`
	ListOrder   *float64 `json:"listOrder,optional"`
	Status      *int     `json:"status,optional"`
}

type FormDatasReq struct {
	FormId   string `form:"formId"`
	Current  *int   `form:"current,optional"`
	PageSize *int   `form:"pageSize,optional"`
}

type FormSearchReq struct {
	FormId          string `json:"formId"`
	SearchFieldJson string `json:"searchFieldJson,optional"`
	Current         *int   `json:"current,optional"`
	PageSize        *int   `json:"pageSize,optional"`
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

type SettingShowReq struct {
	Key string `path:"key"`
}

type SettingSaveReq struct {
	Key          string `json:"key"`
	FormDataJson string `json:"formDataJson"`
}

type AdminMenuGetReq struct {
	Plugin string `form:"plugin,optional"`
}

type AdminMenuShowReq struct {
	FormId string `path:"formId"`
}

type AdminMenuSaveReq struct {
	Id          string   `json:"id,optional"`
	ParentId    *string  `json:"parentId,optional"`
	FormId      *string  `json:"formId,optional"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon,optional"`
	MenuType    int      `json:"menuType,optional"`
	HideInMenu  *int     `json:"hideInMenu,optional"`
	Description string   `json:"description,optional"`
	ListOrder   *float64 `json:"listOrder,optional"`
	Status      *int     `json:"status,optional"`
}

type ThemeListReq struct {
	Current  *int `form:"current,optional"`
	PageSize *int `form:"pageSize,optional"`
}

type ThemeShowReq struct {
	Id int `json:"id"`
}

type ThemeSaveReq struct {
	Id          int    `path:"id,optional"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
	Version     string `json:"version"`
}

type ThemePageListReq struct {
	ThemeKey string  `path:"themeKey"`
	Name     *string `form:"name,optional"`
	Type     string  `form:"type,optional"`
	IsPublic *int    `form:"isPublic,optional"`
	Current  *int    `form:"current,optional"`
	PageSize *int    `form:"pageSize,optional"`
	Status   *int    `form:"status,optional"`
}

type ThemePageShowReq struct {
	Id       string `path:"id"`
	ThemeKey string `form:"themeKey,optional"`
	Type     string `form:"type,optional"`
}

type ThemePageSaveReq struct {
	Id             string  `path:"id,optional"`
	ThemeKey       string  `json:"themeKey"`
	IsPublic       int     `json:"isPublic,optional"`
	Name           string  `json:"name"`
	Alias          string  `json:"alias,optional"`
	Description    string  `json:"description,optional"`
	Schema         string  `json:"schema,optional"`
	SeoTitle       string  `json:"seoTitle,optional"`
	SeoKeywords    string  `json:"seoKeywords,optional"`
	SeoDescription string  `json:"seoDescription,optional"`
	Type           string  `json:"type"`
	ListOrder      float64 `json:"listOrder,optional"`
	Status         *int    `json:"status,optional"`
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
