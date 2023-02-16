// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	useradminaccount "zerocmf/service/user/api/internal/handler/user/admin/account"
	useradminauthAccess "zerocmf/service/user/api/internal/handler/user/admin/authAccess"
	useradminauthorize "zerocmf/service/user/api/internal/handler/user/admin/authorize"
	useradmindepartment "zerocmf/service/user/api/internal/handler/user/admin/department"
	useradminrole "zerocmf/service/user/api/internal/handler/user/admin/role"
	userapp "zerocmf/service/user/api/internal/handler/user/app"
	useroauth "zerocmf/service/user/api/internal/handler/user/oauth"
	"zerocmf/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: useradminaccount.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: useradminaccount.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: useradminaccount.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: useradminaccount.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: useradminaccount.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/account"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: useradminrole.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: useradminrole.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: useradminrole.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: useradminauthorize.GetHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/authorize"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: useradminauthAccess.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: useradminauthAccess.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: useradminauthAccess.EditHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/auth_access"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: userapp.IndexHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/save",
				Handler: userapp.SaveHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/current_user",
					Handler: userapp.CurrentUserHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token",
				Handler: useroauth.TokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/refresh",
				Handler: useroauth.RefreshHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/token",
				Handler: useroauth.TokenRequestHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/validation/token",
				Handler: useroauth.ValidationHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: useradmindepartment.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: useradmindepartment.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: useradmindepartment.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: useradmindepartment.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: useradmindepartment.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/department"),
	)
}
