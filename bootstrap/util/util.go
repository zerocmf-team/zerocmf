/**
** @创建时间: 2021/11/27 08:46
** @作者　　: return
** @描述　　:
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/db"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

//GetMd5 获取加盐的md5字符串

func GetMd5(s string) string {
	conf := config.Config()
	h := md5.New()
	h.Write([]byte(conf.Database.AuthCode + s))
	return hex.EncodeToString(h.Sum(nil))
}

func TenantId(c *gin.Context) string {
	iTenantId, exist := c.Get("tenant_id")
	tenantId := ""
	if exist {
		switch iTenantId.(type) {
		case int:
			idInt := iTenantId.(int)
			if idInt > 0 {
				tenantId = strconv.Itoa(iTenantId.(int))
			}
		case string:
			tenantId = iTenantId.(string)
		}
	}
	return tenantId
}

func ManualDb(tenantId string) *gorm.DB {
	return db.ManualDb(tenantId)
}

func GetDb(c *gin.Context) *gorm.DB {
	iTenantId, exist := c.Get("tenant_id")
	tenantId := ""
	if exist {
		switch iTenantId.(type) {
		case int:
			idInt := iTenantId.(int)
			if idInt > 0 {
				tenantId = strconv.Itoa(iTenantId.(int))
			}
		case string:
			tenantId = iTenantId.(string)
		}
	}
	return db.ManualDb(tenantId)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description gorm数据库链接是否异常，排除空错
 * @Date 2021/12/6 14:27:55
 * @Param
 * @return
 **/

func IsDbErr(db *gorm.DB) error {
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return db.Error
	}
	return nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取文件的预览url地址
 * @Date 2021/12/6 14:34:2
 * @Param
 * @return
 **/

func FileUrl(path string) string {
	conf := config.Config()
	if path == "" {
		return ""
	}
	domain := conf.App.Domain
	prevPath := domain + "/public/uploads/" + path
	return prevPath
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 是否在数组中
 * @Date 2021/12/8 11:34:53
 * @Param
 * @return
 **/

func ToLowerInArray(search string, target []string) bool {
	for _, item := range target {
		if strings.ToLower(search) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
