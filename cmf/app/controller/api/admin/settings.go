/**
** @创建时间: 2021/12/1 12:51
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"encoding/json"
	"gincmf/app/model"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"gorm.io/gorm"
)

type Settings struct {
	controller.Rest
}

func (rest *Settings) Get(c *gin.Context) {
	db := util.GetDb(c)
	option := &model.Option{}
	tx := db.Where("option_name = ?", "site_info").First(option) // 查询
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		rest.Error(c, "获取失败："+tx.Error.Error(), nil)
		return
	}
	rest.Success(c, "获取成功", option)
}

func (rest *Settings) Store(c *gin.Context) {

	var form model.SiteInfo
	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c,"参数有误："+err.Error(),nil)
		return
	}

	siteInfoValue, _ := json.Marshal(form)
	db := util.GetDb(c)
	tx := db.Model(&model.Option{}).Where("option_name = ?", "site_info").Update("option_value", string(siteInfoValue))
	if tx.Error != nil {
		rest.Error(c, "系统出错："+tx.Error.Error(), nil)
		return
	}
	rest.Success(c, "修改成功", form)
}
