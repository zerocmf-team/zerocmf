package role

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/admin/role"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewShowLogic(r.Context(), svcCtx)
		resp, _ := l.Show(&req)
		httpx.OkJson(w, resp)
	}
}
