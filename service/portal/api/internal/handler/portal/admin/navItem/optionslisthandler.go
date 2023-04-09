package navItem

import (
	"net/http"
	"zerocmf/service/portal/api/internal/logic/portal/admin/navItem"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func OptionsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NavItemOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := navItem.NewOptionsListLogic(r.Context(), svcCtx)
		resp := l.OptionsList(&req)
		httpx.OkJson(w, resp)
	}
}
