/**
** @创建时间: 2022/2/27 09:30
** @作者　　: return
** @描述　　:
 */

package middleware

import "github.com/gin-gonic/gin"

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
