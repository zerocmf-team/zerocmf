/**
** @创建时间: 2021/11/28 11:19
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"gincmf/app/grpc/oauth"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"strings"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 验证access_token是否有效
 * @Date 2021/11/28 11:19:35
 * @Param
 * @return
 **/

func ValidationBearerToken(c *gin.Context) {

	r := c.Request

	auth := r.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = r.FormValue("access_token")
	}

	if token != "" {

		userId, err := new(oauth.Oauth).Request(token)

		if err != nil {
			new(controller.Rest).Error(c, err.Error(), nil)
			c.Abort()
			return
		}

		if userId != "" {
			c.Set("userId", userId)
		}
	}

	c.Next()
}
