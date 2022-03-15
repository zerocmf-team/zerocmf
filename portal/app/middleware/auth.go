/**
** @创建时间: 2021/12/10 17:27
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

func ValidationBearerToken(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		c.Abort()
		c.JSON(401,gin.H{
			"code":"0",
			"msg":"用户id不能为空！",
		})
		return
	}
	c.Set("userId",userId)
	c.Next()
}
