/**
** @创建时间: 2021/11/23 09:46
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/controller/api/app"
	web "gincmf/app/controller/app"
	"github.com/gin-gonic/gin"
	bsMidWare "github.com/gincmf/bootstrap/middleware"
)

func ApiListen(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	e.StaticFS("/public", gin.Dir("public/", true))
	v1 := e.Group("api/v1")
	v1.Use(bsMidWare.Init)
	{
		v1.GET("/", new(admin.Index).Get)
		adminGroup := v1.Group("/admin")
		{
			adminGroup.Use(bsMidWare.ValidationBearerToken)
			adminGroup.GET("/admin_menu", new(admin.Menu).Get)
			adminGroup.GET("/settings", new(admin.Settings).Get)
			adminGroup.POST("/settings", new(admin.Settings).Store)
			adminGroup.GET("/upload", new(admin.Upload).Get)
			adminGroup.GET("/assets", new(admin.Assets).Get)
			adminGroup.POST("/assets", new(admin.Assets).Store)
		}

		appGroup := v1.Group("/app")
		{
			appGroup.GET("/route", new(web.Route).List)
			appGroup.GET("/comment/:id",new(app.Comment).Get)
			appToken := appGroup.Use(bsMidWare.ValidationBearerToken)
			appToken.POST("/comment/:id", new(app.Comment).Comment)
			appToken.POST("/comment/like/:id", new(app.Comment).Like)
			appToken.POST("/comment/reply/:id", new(app.Comment).Reply)
			appToken.POST("/comment/reply/like/:id", new(app.Comment).ReplyLike)
		}
	}
}
