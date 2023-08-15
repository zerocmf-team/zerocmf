/**
** @创建时间: 2022/3/13 15:58
** @作者　　: return
** @描述　　:
 */

package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/Init"
)

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(Init.Salts() + s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
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
	prevPath := Init.Domain() + "/public/uploads/" + path
	return prevPath
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 是否在数组中
 * @Date 2021/12/8 11:34:53
 * @Param
 * @return
 **/

/*func ToLowerInArray(search string, target []string) bool {
	for _, item := range target {
		if strings.ToLower(search) == strings.ToLower(item) {
			return true
		}
	}
	return false
}*/

func ToLowerInArray(search string, target interface{}) bool {

	switch target.(type) {
	case []int:
		targetArr := target.([]int)
		for _, item := range targetArr {
			itemStr := strconv.Itoa(item)
			if strings.ToLower(search) == strings.ToLower(itemStr) {
				return true
			}
		}
	case []string:
		targetArr := target.([]string)
		for _, item := range targetArr {
			if strings.ToLower(search) == strings.ToLower(item) {
				return true
			}
		}
	}

	return false
}

func Host(r *http.Request) (domain string) {
	// 获取请求头域名
	scheme := "http://"
	if r.Header.Get("Scheme") == "https" {
		scheme = "https://"
	}
	host := r.Host
	domain = scheme + host
	return
}

func RemoveDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueArr := []string{}

	for _, str := range arr {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			uniqueArr = append(uniqueArr, str)
		}
	}

	return uniqueArr
}
