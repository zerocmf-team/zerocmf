// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	adminMenu "zerocmf/service/admin/api/internal/handler/adminMenu"
	assets "zerocmf/service/admin/api/internal/handler/assets"
	optionadmin "zerocmf/service/admin/api/internal/handler/option/admin"
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
					Method:  http.MethodPost,
					Path:    "/admin_menu/sync",
					Handler: adminMenu.SyncHandler(serverCtx),
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
					Handler: optionadmin.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/settings",
					Handler: optionadmin.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/upload",
					Handler: optionadmin.UploadGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: optionadmin.UploadStoreHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1"),
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