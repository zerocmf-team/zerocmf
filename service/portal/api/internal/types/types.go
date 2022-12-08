// Code generated by goctl. DO NOT EDIT.
package types

import (
	"github.com/jinzhu/copier"
	bsData "zerocmf/common/bootstrap/data"
)

type OneReq struct {
	Id int `path:"id,optional"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CateGetReq struct {
	Name string `form:"name,optional"`
}

type CateSaveReq struct {
	Id             int    `path:"id,optional"`
	Name           string `json:"name,optional"`
	ParentId       int    `json:"parent_id,optional"`
	Status         int    `json:"status,optional"`
	Alias          string `json:"alias,optional"`
	Description    string `json:"description,optional"`
	Thumbnail      string `json:"thumbnail,optional"`
	SeoTitle       string `json:"seo_title,optional"`
	SeoDescription string `json:"seo_description,optional"`
	SeoKeywords    string `json:"seo_keywords,optional"`
	ListTpl        string `json:"list_tpl,optional"`
	OneTpl         string `json:"one_tpl,optional"`
}

type CateDelReq struct {
	Id int `path:"id,optional"`
}

type ArticleGetReq struct {
	Title     string `form:"post_title,optional"`
	PostType  string `form:"post_type,optional"`
	StartTime string `form:"start_time,optional"`
	EndTime   string `form:"end_time,optional"`
}

type ArticleSearchReq struct {
	Keywords string `form:"keywords,optional"`
}

type Extends struct {
	Key   string `json:"key,optional"`
	Value string `json:"value,optional"`
}

type Path struct {
	RemarkName string `json:"remark_name,optional"`
	PrevPath   string `json:"prev_path,optional"`
	FilePath   string `json:"file_path,optional"`
}

type ArticleSaveReq struct {
	Id             int       `path:"id,optional"`
	CategoryIds    []string  `json:"category_ids,optional"`
	PostType       int       `json:"post_type,optional"`
	Alias          string    `json:"alias,optional"`
	PostTitle      string    `json:"post_title,optional"`
	Thumbnail      string    `json:"thumbnail,optional"`
	PostKeywords   []string  `json:"post_keywords,optional"`
	ListOrder      float64   `json:"list_order,optional"`
	PublishTime    string    `json:"publish_time,optional"`
	PostSource     string    `json:"post_source,optional"`
	PostExcerpt    string    `json:"post_excerpt,optional"`
	PostContent    string    `json:"post_content,optional"`
	IsTop          int       `json:"is_top,optional"`
	SeoTitle       string    `json:"seo_title,optional"`
	SeoKeywords    string    `json:"seo_keywords,optional"`
	SeoDescription string    `json:"seo_description,optional"`
	Recommended    int       `json:"recommended,optional"`
	PostStatus     int       `json:"post_status,optional"`
	Photos         []Path    `json:"photos,optional"`
	Files          []Path    `json:"files,optional"`
	Audio          string    `json:"audio,optional"`
	Video          string    `json:"video,optional"`
	Template       string    `json:"template,optional"`
	Extends        []Extends `json:"extends,optional"`
}

type ArticleDelReq struct {
	Id int `path:"id,optional"`
}

type PostListReq struct {
	Ids string `form:"ids,optional"`
	Hot int    `form:"hot,optional"` // 根据浏览量和权重排序
}

type PostShowReq struct {
	Id   int `path:"id,optional"`
	Type int `form:"type,optional"`
}

type InitReq struct {
	Theme     string   `json:"theme,optional"`
	Version   string   `json:"version,optional"`
	Thumbnail string   `json:"thumbnail,optional"`
	ThemeFile []string `json:"theme_file,optional"`
}

type ListReq struct {
	Type string `form:"type,optional"`
}

type ThemeFileSaveReq struct {
	Id   string `path:"id"`
	More string `json:"more,optional"`
}

type ThemeFileListReq struct {
	Theme    string `form:"theme,optional"`
	IsPublic string `form:"is_public,optional"`
}

type ThemeFileDetailReq struct {
	Theme string `form:"theme,optional"`
	File  string `form:"file,optional"`
}

type TagGetReq struct {
	Name string `form:"name,optional"`
}

type NavItemGetReq struct {
	Key string `form:"key,optional"`
}

type NavItemSaveReq struct {
	Id        int     `path:"id,optional"`
	NavId     int     `json:"nav_id,optional"`
	ParentId  int     `json:"parent_id,optional"`
	Status    int     `json:"status,optional"`
	ListOrder float64 `json:"list_order,optional"`
	Name      string  `json:"name,optional"`
	Target    string  `json:"target,optional"`
	HrefType  int     `json:"href_type,optional"`
	Href      string  `json:"href,optional"`
	Icon      string  `json:"icon,optional"`
	Path      string  `json:"path,optional"`
}

type NavItemOptionsReq struct {
	NavId int `form:"nav_id"`
}

type PostCommentGetReq struct {
	Id   int `path:"id,optional"`
	Type int `form:"type,optional"`
}

type PostCommentAddReq struct {
	Id        int    `path:"id,optional"`
	TopicType int    `json:"topic_type,optional"`
	Content   string `json:"content,optional"`
}

type PostReplyReq struct {
	Id        int    `path:"id,optional"`
	ReplyId   int    `json:"reply_id,optional"`
	ReplyType int    `json:"reply_type,optional"`
	Content   string `json:"content,optional"`
	ToUserId  int    `json:"to_user_id,optional"`
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
