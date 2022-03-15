/**
** @创建时间: 2022/3/5 14:16
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 自动修正domain
 * @Date 2021/12/6 21:19:15
 * @Param
 * @return
 **/

func Init(c *gin.Context) {
	conf := config.Config()
	if conf.App.Domain == "" {
		scheme := "http://"
		if c.Request.Header.Get("Scheme") == "https" {
			scheme = "https://"
		}
		domain := scheme + c.Request.Host
		config.SetDomain(domain)
	}
}
