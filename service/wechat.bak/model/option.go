/**
** @创建时间: 2021/11/26 10:33
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/util"
)

type Option struct {
	Id          int    `json:"id"`
	AutoLoad    int    `gorm:"type:tinyint(3);default:1;not null" json:"autoload"`
	OptionName  string `gorm:"type:varchar(64);not null" json:"option_name"`
	OptionValue string `gorm:"type:longtext" json:"option_value"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 定义微信小程序渠道
 * @Date 2022/4/22 9:46:13
 * @Param
 * @return
 **/

type MiniProgram struct {
	AppID        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	OriginalId   string `json:"original_id"`   // 原始id
	UploadTicket string `json:"upload_ticket"` // 小程序上陈秘钥
}

type MediaPlatform struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	OriginalId string `json:"original_id"` // 原始id
}

func (_ *Option) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Option{})
	// 初始化小程序配置项
	tx := db.Where("option_name = ?", "wxapp").First(&Option{})
	if tx.RowsAffected == 0 {
		mp := MiniProgram{}
		ov, _ := json.Marshal(mp)
		db.Create(&Option{AutoLoad: 1, OptionName: "wxapp", OptionValue: string(ov)})
	}

	// 初始化公众号配置项
	tx = db.Where("option_name = ?", "mp").First(&Option{}) // 查询
	if tx.RowsAffected == 0 {
		//初始化默认json
		mp := MediaPlatform{}
		ov, _ := json.Marshal(mp)
		db.Create(&Option{AutoLoad: 1, OptionName: "mp", OptionValue: string(ov)})
	}
}

func (model *MiniProgram) Show(db *gorm.DB) (res MiniProgram, err error) {
	op := Option{}
	tx := db.Where("option_name = ?", "wxapp").First(&op)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	json.Unmarshal([]byte(op.OptionValue), &res)
	return
}

func (model *MediaPlatform) Show(db *gorm.DB) (res MediaPlatform, err error) {
	op := Option{}
	tx := db.Where("option_name = ?", "mp").First(&op)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}
	json.Unmarshal([]byte(op.OptionValue), &res)
	return
}
