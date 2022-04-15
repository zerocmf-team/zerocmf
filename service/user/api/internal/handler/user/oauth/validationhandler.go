package oauth

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/oauth"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ValidationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewValidationLogic(r.Context(), svcCtx)
		resp, _ := l.Validation(&req)
		httpx.OkJson(w, resp)
	}
}
