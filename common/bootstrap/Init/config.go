/**
** @创建时间: 2022/3/13 16:55
** @作者　　: return
** @描述　　:
 */

package Init

import (
	"sync"
)

var (
	mu     sync.RWMutex
	domain string
	salts  string
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 设置域名
 * @Date 2022/3/13 16:57:18
 * @Param
 * @return
 **/

func SetDomain(value string) {
	mu.Lock()
	domain = value
	mu.Unlock()
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取域名
 * @Date 2022/3/13 16:57:27
 * @Param
 * @return
 **/

func Domain() string {
	return domain
}

func SetSalts(value string) {
	mu.Lock()
	salts = value
	mu.Unlock()
}

func Salts() string {
	return salts
}
