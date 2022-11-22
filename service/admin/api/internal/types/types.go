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

type MenuReq struct {
	Id         int     `path:"id,optional"`
	ParentId   int     `json:"parent_id"`
	MenuType   int     `json:"menu_type" validate:"oneof=0 1" label:"菜单类型"`
	Name       string  `json:"name" validate:"required" label:"菜单名称"`
	Path       string  `json:"path" validate:"required" label:"路由地址"`
	Icon       string  `json:"icon"`
	HideInMenu int     `json:"hide_in_menu"`
	ListOrder  float64 `json:"list_order"`
}

type AssetsReq struct {
	Type string `form:"type,optional,default=0"`
}

type DeleteReq struct {
	Id  int   `path:"id,optional"`
	Ids []int `json:"ids,optional"`
}

type OptionReq struct {
	SiteName           string `json:"site_name,optional"`
	AdminPassword      string `json:"admin_password,optional"`
	SiteSeoTitle       string `json:"site_seo_title,optional"`
	SiteSeoKeywords    string `json:"site_seo_keywords,optional"`
	SiteSeoDescription string `json:"site_seo_description,optional"`
	SiteIcp            string `json:"site_icp,optional"`
	SiteGwa            string `json:"site_gwa,optional"`
	SiteAdminEmail     string `json:"site_admin_email,optional"`
	SiteAnalytics      string `json:"site_analytics,optional"`
	OpenRegistration   int    `json:"open_registration,optional"`
}

type Assets struct {
	UploadMaxFileSize int    `json:"upload_max_file_size,optional"`
	Extensions        string `json:"extensions,optional"`
}

type FileTypes struct {
	Image Assets `json:"image,optional"`
	Video Assets `json:"video,optional"`
	Audio Assets `json:"audio,optional"`
	File  Assets `json:"file,optional"`
}

type UploadReq struct {
	MaxFiles  int       `json:"max_files,optional"`
	ChunkSize int       `json:"chunk_size,optional"`
	FileTypes FileTypes `json:"file_types,optional"`
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
