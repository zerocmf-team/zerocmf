package navItem

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/navItem"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := navItem.NewDelLogic(r.Context(), svcCtx)
		resp := l.Del(&req)
		httpx.OkJson(w, resp)
	}
}
