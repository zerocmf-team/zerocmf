package navItem

import (
	"net/http"
	"zerocmf/service/portal/api/internal/logic/portal/admin/navItem"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NavItemSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := navItem.NewEditLogic(r.Context(), svcCtx)
		resp := l.Edit(&req)
		httpx.OkJson(w, resp)
	}
}
