/**
** @创建时间: 2021/12/2 12:52
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

type Upload struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取上传设置信息
 * @Date 2021/12/2 12:59:15
 * @Param
 * @return
 **/

func (rest *Upload) Get(c *gin.Context)  {

	db := util.GetDb(c)
	option := model.Option{}
	uploadSetting := model.UploadSetting{}
	tx := db.Where("option_name = ?", "upload_setting").First(&option)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		rest.Error(c,"系统出错："+tx.Error.Error(),nil)
		return
	}
	value := option.OptionValue
	err := json.Unmarshal([]byte(value), &uploadSetting)

	if err != nil {
		rest.Error(c,"解析时出错："+err.Error(),nil)
		return
	}

	rest.Success(c,"获取成功！",uploadSetting)

}
