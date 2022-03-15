/**
** @创建时间: 2021/11/23 09:46
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/middleware"
	"github.com/gin-gonic/gin"
)

func ApiListen(e *gin.Engine) {
	v1 := e.Group("api/v1")
	v1.Use(middleware.Init)
	{
		v1.GET("/", new(admin.Index).Get)
		adminGroup := v1.Group("/admin")
		adminGroup.Use(middleware.ValidationBearerToken)
		adminGroup.GET("/index", new(admin.Index).Get)
	}
}
