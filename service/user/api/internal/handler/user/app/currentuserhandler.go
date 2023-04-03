package app

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/user/api/internal/logic/user/app"
	"zerocmf/service/user/api/internal/svc"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := app.NewCurrentUserLogic(r.Context(), svcCtx)
		resp := l.CurrentUser()
		httpx.OkJson(w, resp)
	}
}
