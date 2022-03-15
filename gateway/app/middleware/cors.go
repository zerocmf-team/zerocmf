/**
** @创建时间: 2020/11/1 10:34 上午
** @作者　　: return
** @描述　　:
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"regexp"
)

func AllowCors(c *gin.Context) {

	origin := c.Request.Header.Get("Origin")
	var filterHost = [...]string{}
	// filterHost 做过滤器，防止不合法的域名访问
	var isAccess = false
	for _, v := range filterHost {
		match, _ := regexp.MatchString(v, origin)
		if match {
			isAccess = true
		}
	}

	if len(filterHost) == 0 {
		isAccess = true
	}

	if isAccess {
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
	}

	//放行所有OPTIONS方法
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.Next()
}
