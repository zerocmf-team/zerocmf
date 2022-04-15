package app

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/app"
	"gincmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := app.NewCurrentUserLogic(r.Context(), svcCtx)
		resp, _ := l.CurrentUser()
		httpx.OkJson(w, resp)
	}
}
