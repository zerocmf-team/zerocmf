/**
** @创建时间: 2020/8/31 11:37 上午
** @作者　　: return
** @描述　　:
 */
package oauth

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestRegisterOauthRouter(t *testing.T) {
}
