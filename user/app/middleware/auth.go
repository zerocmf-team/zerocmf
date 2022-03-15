/**
** @创建时间: 2021/11/28 11:19
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/middleware"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 验证access_token是否有效
 * @Date 2021/11/28 11:19:35
 * @Param
 * @return
 **/

func ValidationBearerToken(c *gin.Context) {
	middleware.ValidationBearerToken(c)
}


