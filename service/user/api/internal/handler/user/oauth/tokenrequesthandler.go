package oauth

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/oauth"
	"gincmf/service/user/api/internal/svc"
)

func TokenRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewTokenRequestLogic(r.Context(), svcCtx)
		l.TokenRequest()
		//resp, _ := l.TokenRequest()
		//httpx.OkJson(w, resp)
	}
}
