/**
** @创建时间: 2021/11/26 10:33
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/util"
)

type Option struct {
	Id          int    `json:"id"`
	AutoLoad    int    `gorm:"type:tinyint(3);default:1;not null" json:"autoload"`
	OptionName  string `gorm:"type:varchar(64);not null" json:"option_name"`
	OptionValue string `gorm:"type:longtext" json:"option_value"`
}

//定义site_info类型

type SiteInfo struct {
	SiteName           string `json:"site_name"`
	AdminPassword      string `json:"admin_password"`
	SiteSeoTitle       string `json:"site_seo_title"`
	SiteSeoKeywords    string `json:"site_seo_keywords"`
	SiteSeoDescription string `json:"site_seo_description"`
	SiteIcp            string `json:"site_icp"`
	SiteGwa            string `json:"site_gwa"`
	SiteAdminEmail     string `json:"site_admin_email"`
	SiteAnalytics      string `json:"site_analytics"`
	OpenRegistration   int    `json:"open_registration"`
}

//定义upload_setting类型

type UploadSetting struct {
	MaxFiles  int `json:"max_files"`
	ChunkSize int `json:"chunk_size"`
	FileTypes `json:"file_types"`
}

type FileTypes struct {
	Image TypeValues `json:"image"`
	Video TypeValues `json:"video"`
	Audio TypeValues `json:"audio"`
	File  TypeValues `json:"file"`
}

type TypeValues struct {
	UploadMaxFileSize int    `json:"upload_max_file_size"`
	Extensions        string `json:"extensions"`
}

// 手机号登录设置

type MobileLoginSettings struct {
	Platform        string `json:"platform"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SignName        string `json:"sign_name"`
	TemplateCode    string `json:"template_code"`
	TemplateParam   string `json:"template_param"`
	Status          int    `json:"status"`
}

// 微信小程序登录设置

type WxappLoginSettings struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
	Token     string `json:"token"`
	Status    int    `json:"status"`
}

func (_ *Option) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Option{})

	tx := db.Where("option_name = ?", "site_info").First(&Option{})
	if tx.RowsAffected == 0 {
		siteInfo := SiteInfo{}
		siteInfoValue, _ := json.Marshal(siteInfo)
		db.Create(&Option{AutoLoad: 1, OptionName: "site_info", OptionValue: string(siteInfoValue)})
	}

	tx = db.Where("option_name = ?", "upload_setting").First(&Option{}) // 查询
	if tx.RowsAffected == 0 {
		//初始化默认json
		uploadSetting := UploadSetting{
			MaxFiles:  20,
			ChunkSize: 512,
			FileTypes: FileTypes{
				Image: TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "jpg,jpeg,png,gif,bmp4,svg",
				},
				Video: TypeValues{
					UploadMaxFileSize: 102400,
					Extensions:        "mp4,avi,wmv,rm,rmvb,mkv",
				},
				Audio: TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "mp3,wma,wav",
				},
				File: TypeValues{
					UploadMaxFileSize: 10240,
					Extensions:        "txt,pdf,doc,docx,xls,xlsx,ppt,pptx,zip,rar",
				},
			},
		}
		uploadSettingValue, _ := json.Marshal(uploadSetting)
		db.Create(&Option{AutoLoad: 1, OptionName: "upload_setting", OptionValue: string(uploadSettingValue)})
	}

	tx = db.Where("option_name = ?", "mobile_login_setting").First(&Option{})
	if tx.RowsAffected == 0 {
		mobileLogin := MobileLoginSettings{
			Status: 0,
		}
		mobileLoginValue, _ := json.Marshal(mobileLogin)
		db.Create(&Option{AutoLoad: 1, OptionName: "mobile_login_setting", OptionValue: string(mobileLoginValue)})
	}

	tx = db.Where("option_name = ?", "wxapp_login_setting").First(&Option{})
	if tx.RowsAffected == 0 {
		wxappLogin := WxappLoginSettings{
			Status: 0,
		}
		wxappLoginValue, _ := json.Marshal(wxappLogin)
		db.Create(&Option{AutoLoad: 1, OptionName: "wxapp_login_setting", OptionValue: string(wxappLoginValue)})
	}
}

func UploadSettings(db *gorm.DB) (uploadSetting UploadSetting, err error) {
	option := Option{}
	tx := db.Where("option_name = ?", "upload_setting").First(&option) // 查询
	if err := util.IsDbErr(tx); err != nil {
		return uploadSetting, errors.New("数据库出现问题：" + err.Error())
	}
	err = json.Unmarshal([]byte(option.OptionValue), &uploadSetting)
	if err != nil {
		return uploadSetting, errors.New("序列化json出错：" + err.Error())
	}
	return
}
