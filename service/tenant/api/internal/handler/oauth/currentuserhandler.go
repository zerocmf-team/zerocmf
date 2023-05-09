package oauth

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/tenant/api/internal/logic/oauth"
	"zerocmf/service/tenant/api/internal/svc"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewCurrentUserLogic(r.Context(), svcCtx)
		resp := l.CurrentUser()
		httpx.OkJson(w, resp)
	}
}
