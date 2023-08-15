package nav

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/nav"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NavSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := nav.NewStoreLogic(r, svcCtx)
		resp := l.Store(&req)
		httpx.OkJson(w, resp)
	}
}
