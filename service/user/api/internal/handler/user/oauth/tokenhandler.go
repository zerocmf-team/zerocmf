package oauth

import (
	"gincmf/service/user/api/internal/types"
	"net/http"

	"gincmf/service/user/api/internal/logic/user/oauth"
	"gincmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.TokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewTokenLogic(r.Context(), svcCtx)
		resp, _ := l.Token(&req)
		httpx.OkJson(w, resp)
	}
}
