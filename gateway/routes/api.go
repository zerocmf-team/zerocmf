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
	bsMidWare "github.com/gincmf/bootstrap/middleware"
)

func ApiListen(e *gin.Engine) {
	e.Use(bsMidWare.Init,middleware.AllowCors)
	e.Any("/api/",new(admin.Index).Get)
	v1 := e.Group("api/v1")
	adminGroup := v1.Group("/admin")
	{
		adminGroup.Use(middleware.ValidationBearerToken,middleware.Rbac)
		adminGroup.Any("/*name", new(admin.Gateway).Register)
	}

	appGroup := v1.Group("/app")
	{
		appGroup.Use(middleware.ValidationBearerToken)
		appGroup.Any("/*name", new(admin.Gateway).Register)
	}

	e.Any("/api/oauth/*name",new(admin.Gateway).Register)
	e.Any("/public/*name", new(admin.Gateway).Register)
}
