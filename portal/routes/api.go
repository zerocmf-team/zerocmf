/**
** @创建时间: 2021/11/23 09:46
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/controller/api/app"
	"gincmf/app/middleware"
	"github.com/gin-gonic/gin"
)

func ApiListen(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "", "pong")
	})
	v1 := e.Group("api/v1")
	v1.Use(middleware.Init)
	{
		v1.GET("/", new(admin.Index).Get)
		adminGroup := v1.Group("/admin/portal")
		{
			adminGroup.Use(middleware.ValidationBearerToken)
			adminGroup.GET("/category", new(admin.PortalCategory).Get)
			adminGroup.POST("/category", new(admin.PortalCategory).Store)
			adminGroup.GET("/category/:id", new(admin.PortalCategory).Show)
			adminGroup.POST("/category/:id", new(admin.PortalCategory).Edit)
			adminGroup.GET("/category_list", new(admin.PortalCategory).List)
			adminGroup.GET("/category_options", new(admin.PortalCategory).Options)
			adminGroup.GET("/theme_file/list", new(admin.ThemeFile).List)
			adminGroup.POST("/theme_file/:id", new(admin.ThemeFile).Save)
			adminGroup.GET("/article", new(admin.PortalPost).Get)
			adminGroup.POST("/article", new(admin.PortalPost).Store)
			adminGroup.GET("/article/:id", new(admin.PortalPost).Show)
			adminGroup.POST("/article/:id", new(admin.PortalPost).Edit)
			adminGroup.GET("/tag", new(admin.Tag).Get)
			adminGroup.DELETE("/tag/:id", new(admin.Tag).Delete)
			adminGroup.POST("/theme", new(admin.Theme).Init)
			adminGroup.POST("/nav_items", new(admin.NavItem).GetNavItems)
			adminGroup.GET("/nav_item_options", new(admin.NavItem).OptionsList)
			adminGroup.GET("/nav_item_urls", new(admin.NavItem).OptionsUrls)
			adminGroup.POST("/nav_item", new(admin.NavItem).Store)
			adminGroup.POST("/nav_item/:id", new(admin.NavItem).Edit)
			adminGroup.DELETE("/nav_item/:id", new(admin.NavItem).Delete)

		}
	}

	appGroup := v1.Group("/app/portal")
	appGroup.GET("/post/:id", new(app.Post).Show)
	appGroup.POST("/list/cid", new(app.Post).ListWithCid)
	appGroup.GET("/breadcrumb/:cid", new(app.PortalCategory).Breadcrumb)
	appGroup.GET("/list/:id", new(app.PortalCategory).Show)
	appGroup.GET("/category/trees/:cid", new(app.PortalCategory).TreeList)
	appGroup.GET("/theme_file", new(app.ThemeFile).Detail)
	appGroup.GET("/theme_files", new(app.ThemeFile).List)
	appGroup.GET("/nav_items/:id", new(app.NavItem).GetNavItems)

	appToken := appGroup.Use(middleware.ValidationBearerToken)
	appToken.GET("/post/like/:id", new(app.Post).Like)
	appToken.GET("/post/is_like/:id", new(app.Post).IsLike)

	appToken.GET("/post/favorite/:id", new(app.Post).Favorite)
	appToken.GET("/post/is_favorite/:id", new(app.Post).IsFavorite)

}
