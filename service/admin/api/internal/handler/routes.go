// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	adminadminMenu "zerocmf/service/admin/api/internal/handler/admin/adminMenu"
	adminassets "zerocmf/service/admin/api/internal/handler/admin/assets"
	adminoptionlogin "zerocmf/service/admin/api/internal/handler/admin/option/login"
	adminoptionsite "zerocmf/service/admin/api/internal/handler/admin/option/site"
	appoption "zerocmf/service/admin/api/internal/handler/app/option"
	"zerocmf/service/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: IndexHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/admin_menu",
					Handler: adminadminMenu.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/admin_menu/all",
					Handler: adminadminMenu.GetAllMenusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu/sync",
					Handler: adminadminMenu.SyncHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu",
					Handler: adminadminMenu.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu/:id",
					Handler: adminadminMenu.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/admin_menu/:id",
					Handler: adminadminMenu.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/assets",
					Handler: adminassets.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/assets",
					Handler: adminassets.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/assets/:id",
					Handler: adminassets.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/settings",
					Handler: adminoptionsite.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/settings",
					Handler: adminoptionsite.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/upload",
					Handler: adminoptionsite.UploadGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: adminoptionsite.UploadStoreHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/mobile",
					Handler: adminoptionlogin.MobileGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/mobile",
					Handler: adminoptionlogin.MobileStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/wxapp",
					Handler: adminoptionlogin.WxappGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/wxapp",
					Handler: adminoptionlogin.WxappStoreHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/settings"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/settings",
				Handler: appoption.GetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app"),
	)
}
