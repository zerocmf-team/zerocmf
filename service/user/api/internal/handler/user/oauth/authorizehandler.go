package oauth

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/oauth"
	"gincmf/service/user/api/internal/svc"
)

func AuthorizeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewAuthorizeLogic(r.Context(), svcCtx)
		l.Authorize()
		//httpx.OkJson(w, resp)
	}
}
