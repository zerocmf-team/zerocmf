package oauth

import (
	"net/http"

	"zerocmf/service/user/api/internal/logic/user/oauth"
	"zerocmf/service/user/api/internal/svc"
)

func TokenRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewTokenRequestLogic(r.Context(), svcCtx)
		l.TokenRequest()
		// resp := l.TokenRequest()
		// httpx.OkJson(w, resp)
	}
}
