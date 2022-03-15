/**
** @创建时间: 2021/1/8 7:28 下午
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"encoding/json"
	"gincmf/app/model"
	appUtil "gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"time"
)

type Theme struct {
	controller.Rest
}

func (rest *Theme) Init(c *gin.Context) {
	var form struct {
		Theme     string   `json:"theme"`
		Version   string   `json:"version"`
		Thumbnail string   `json:"thumbnail"`
		ThemeFile []string `json:"theme_file"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	// 保存主题
	db := util.GetDb(c)
	theme := model.Theme{
		Name:      form.Theme,
		Version:   form.Version,
		Thumbnail: form.Thumbnail,
		CreateAt:  time.Now().Unix(),
	}
	queryTheme := model.Theme{}
	tx := db.Where("name = ?", form.Theme).First(&queryTheme)
	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}
	id := queryTheme.Id
	if id == 0 {
		db.Create(&theme)
	} else {
		theme.Id = id
		theme.UpdateAt = time.Now().Unix()
		db.Save(&theme)
	}
	var addThemeFile []model.ThemeFile
	for _, v := range form.ThemeFile {

		var file struct {
			Name     string `json:"name"`
			File     string `json:"file"`
			Type     string `json:"type"`
			Desc     string `json:"desc"`
			IsPublic int    `json:"is_public"`
		}

		json.Unmarshal([]byte(v), &file)

		queryFile := model.ThemeFile{}
		tx := db.Where("theme = ? AND file = ?", form.Theme, file.File).First(&queryFile)
		if util.IsDbErr(tx) != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		// 合并json
		var m1, m2 map[string]interface{}
		json.Unmarshal([]byte(v), &m1)              // 源文件
		json.Unmarshal([]byte(queryFile.More), &m2) // 数据库
		merged := appUtil.JsonMerge(m1, m2)
		bytes, err := json.Marshal(merged)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}
		more := string(bytes)
		themeFile := model.ThemeFile{
			Theme:       form.Theme,
			IsPublic:    file.IsPublic,
			Name:        file.Name,
			File:        file.File,
			Type:        file.Type,
			Description: file.Desc,
			More:        more,
			ConfigMore:  v,
			CreateAt:    time.Now().Unix(),
		}

		id := queryFile.Id
		if id == 0 {
			addThemeFile = append(addThemeFile, themeFile)
		} else {
			themeFile.Id = id
			themeFile.UpdateAt = time.Now().Unix()
			tx := db.Save(&themeFile)
			if tx.Error != nil {
				rest.Error(c, tx.Error.Error(), nil)
				return
			}
		}
	}
	if len(addThemeFile) > 0 {
		tx = db.Create(&addThemeFile)
		if util.IsDbErr(tx) != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}
	}

	option := model.Option{
		OptionName: "theme",
		OptionValue: form.Theme,
	}
	queryOption :=  model.Option{}
	tx = db.Where("option_name", "theme").First(&queryOption)
	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	if queryOption.Id == 0 {
		tx = db.Create(&option)
	}else {
		option.Id = queryOption.Id
		tx = db.Save(&option)
	}

	if tx.Error != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	rest.Success(c, "操作成功！", form)
}
