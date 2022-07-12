package app

import (
	"net/http"

	"zerocmf/service/user/api/internal/logic/user/app"
	"zerocmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := app.NewIndexLogic(r.Context(), svcCtx)
		resp := l.Index()
		httpx.OkJson(w, resp)
	}
}
