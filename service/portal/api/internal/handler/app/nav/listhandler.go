package nav

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/app/nav"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NavItemGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := nav.NewListLogic(r, svcCtx)
		resp := l.List(&req)
		httpx.OkJson(w, resp)
	}
}
