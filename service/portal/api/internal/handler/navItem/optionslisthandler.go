package navItem

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/navItem"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
