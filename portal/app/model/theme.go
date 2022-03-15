/**
** @创建时间: 2022/1/14 20:09
** @作者　　: return
** @描述　　:
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/gincmf/bootstrap/util"
	"gorm.io/gorm"
)

type Theme struct {
	Id        int     `json:"id"`
	Name      string  `gorm:"type:varchar(40);comment:主题名称;not null" json:"name"`
	Version   string  `gorm:"type:varchar(10);comment:主题版本;not null" json:"version"`
	Thumbnail string  `gorm:"type:varchar(255);comment:主题缩略图;not null" json:"thumbnail"`
	CreateAt  int64   `gorm:"type:int(10);comment:创建时间;default:0" json:"create_at"`
	UpdateAt  int64   `gorm:"type:int(10);comment:更新时间;default:0" json:"update_at"`
	ListOrder float64 `gorm:"type:float;comment:排序;default:10000" json:"list_order"`
	DeleteAt  int64   `gorm:"type:int(10);comment:删除时间;default:0" json:"delete_at"`
}

type ThemeFile struct {
	Id             int         `json:"id"`
	Theme          string      `gorm:"type:varchar(40);comment:主题名称;not null" json:"theme"`
	ListOrder      float64     `gorm:"type:float;comment:排序;default:10000" json:"list_order"`
	IsPublic       int         `gorm:"type:tinyint(3);comment:是否公告部分;not null" json:"is_public"`
	Name           string      `gorm:"type:varchar(20);comment:模板文件名;not null" json:"name"`
	File           string      `gorm:"type:varchar(50);comment:模板文件,对应的魔板渲染页面;not null" json:"file"`
	Type           string      `gorm:"type:varchar(20);comment:模板类型(page：页面，list：列表);default:page;not null" json:"type"`
	Description    string      `gorm:"type:varchar(255);comment:模板描述;not null" json:"description"`
	More           string      `gorm:"type:longtext;comment:主题文件用户配置文件" json:"more"`
	MoreJson       interface{} `gorm:"-" json:"more_json"`
	ConfigMore     string      `gorm:"type:longtext;comment:主题文件默认配置文件" json:"config_more"`
	ConfigMoreJson interface{} `gorm:"-" json:"config_more_json"`
	CreateAt       int64       `gorm:"type:int(11)" json:"create_at"`
	UpdateAt       int64       `gorm:"type:int(11)" json:"update_at"`
}

func (model *Theme) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model)
	db.AutoMigrate(&ThemeFile{})
}

func (model *ThemeFile) Show(db *gorm.DB, query string, queryArgs []interface{}) (ThemeFile, error) {
	themeFile := ThemeFile{}
	result := db.Where(query, queryArgs...).First(&themeFile)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return themeFile, errors.New("该分类不存在！")
		}
		return themeFile, result.Error
	}
	json.Unmarshal([]byte(themeFile.More), &themeFile.MoreJson)
	json.Unmarshal([]byte(themeFile.ConfigMore), &themeFile.ConfigMoreJson)
	return themeFile, nil
}

func (model *ThemeFile) List(db *gorm.DB, query string, queryArgs []interface{}) ([]ThemeFile, error) {

	var themeFile []ThemeFile

	tx := db.Where(query, queryArgs...).Find(&themeFile)

	if util.IsDbErr(tx) != nil {
		return themeFile, tx.Error
	}

	for k, v := range themeFile {
		json.Unmarshal([]byte(v.More), &themeFile[k].MoreJson)
		json.Unmarshal([]byte(v.ConfigMore), &themeFile[k].ConfigMoreJson)
	}

	return themeFile, nil
}
