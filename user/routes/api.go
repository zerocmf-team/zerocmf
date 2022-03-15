/**
** @创建时间: 2021/11/23 09:46
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/controller/api/app"
	"gincmf/app/controller/api/common"
	"gincmf/app/middleware"
	"github.com/gin-gonic/gin"
)

func ApiListen(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	common.Routes(e)
	v1 := e.Group("api/v1")
	v1.Use(middleware.Init)
	{
		v1.GET("/", new(admin.Index).Get)
		appGroup := v1.Group("/app/user")
		appGroup.POST("/register", new(app.User).Register)
		appGroup.GET("/captcha", new(common.Captcha).New)
		appGroup.GET("/captcha/:id", new(common.Captcha).Captcha)
		appGroup.GET("/sms/send", new(common.Sms).Send)
		appGroup.GET("/sms/verify", new(common.Sms).Verify)


		adminGroup := v1.Group("/admin/user")
		adminGroup.Use(middleware.ValidationBearerToken)
		adminGroup.GET("/index", new(admin.Index).Get)
		adminGroup.GET("/current_user", new(admin.User).CurrentUser)
		adminGroup.POST("/save",new(admin.User).Save)
		adminGroup.GET("/account", new(admin.Account).Get)
		adminGroup.GET("/account/:id", new(admin.Account).Show)
		adminGroup.POST("/account/:id", new(admin.Account).Edit)
		adminGroup.POST("/account", new(admin.Account).Store)
		adminGroup.GET("/role", new(admin.Role).Get)
		adminGroup.GET("/role/:id", new(admin.Role).Show)
		adminGroup.DELETE("/role/:id", new(admin.Role).Delete)
		adminGroup.GET("/authorize", new(admin.Authorize).Get)
		adminGroup.POST("/auth_access", new(admin.AuthAccess).Store)
		adminGroup.GET("/auth_access/:id", new(admin.AuthAccess).Show)
		adminGroup.POST("/auth_access/:id", new(admin.AuthAccess).Edit)
	}
}
