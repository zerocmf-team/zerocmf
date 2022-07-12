package authAccess

import (
	"net/http"

	"zerocmf/service/user/api/internal/logic/user/admin/authAccess"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authAccess.NewShowLogic(r.Context(), svcCtx)
		resp := l.Show(&req)
		httpx.OkJson(w, resp)
	}
}
