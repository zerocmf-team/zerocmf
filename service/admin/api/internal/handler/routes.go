// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"zerocmf/service/admin/api/internal/handler/adminMenu"
	"zerocmf/service/admin/api/internal/handler/assets"
	admin2 "zerocmf/service/admin/api/internal/handler/option/admin"
	login2 "zerocmf/service/admin/api/internal/handler/option/admin/login"
	optionapp "zerocmf/service/admin/api/internal/handler/option/app"

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
					Handler: adminMenu.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/admin_menu/all",
					Handler: adminMenu.GetAllMenusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu/sync",
					Handler: adminMenu.SyncHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu",
					Handler: adminMenu.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/admin_menu/:id",
					Handler: adminMenu.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/admin_menu/:id",
					Handler: adminMenu.DeleteHandler(serverCtx),
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
					Handler: assets.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/assets",
					Handler: assets.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/assets/:id",
					Handler: assets.DeleteHandler(serverCtx),
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
					Handler: admin2.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/settings",
					Handler: admin2.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/upload",
					Handler: admin2.UploadGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: admin2.UploadStoreHandler(serverCtx),
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
					Handler: login2.MobileGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/mobile",
					Handler: login2.MobileStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/wxapp",
					Handler: login2.WxappGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/wxapp",
					Handler: login2.WxappStoreHandler(serverCtx),
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
				Handler: optionapp.GetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app"),
	)
}
