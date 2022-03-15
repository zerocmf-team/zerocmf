/**
** @创建时间: 2021/11/23 09:46
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/app/controller/api/admin"
	"github.com/gin-gonic/gin"
)

func ApiListen(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/index", new(admin.Index).Get)
	}
}
