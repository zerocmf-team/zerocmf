package theme

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/admin/theme"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := theme.NewInitLogic(r.Context(), svcCtx)
		resp := l.Init(&req)
		httpx.OkJson(w, resp)
	}
}
