/**
** @创建时间: 2022/3/13 15:58
** @作者　　: return
** @描述　　:
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
	"gincmf/common/bootstrap/data"
	"gorm.io/gorm"
	"strings"
)

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte( data.Salts() + s))
	return hex.EncodeToString(h.Sum(nil))
}

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
	if path == "" {
		return ""
	}
	prevPath := data.Domain() + "/public/uploads/" + path
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
