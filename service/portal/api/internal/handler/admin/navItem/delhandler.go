package navItem

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/navItem"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := navItem.NewDelLogic(r, svcCtx)
		resp := l.Del(&req)
		httpx.OkJson(w, resp)
	}
}
