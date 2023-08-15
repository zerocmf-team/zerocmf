package form

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/form"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FormShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := form.NewShowLogic(r, svcCtx)
		resp := l.Show(&req)
		httpx.OkJson(w, resp)
	}
}
