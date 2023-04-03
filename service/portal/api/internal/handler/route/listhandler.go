package route

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/portal/api/internal/logic/route"
	"zerocmf/service/portal/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := route.NewListLogic(r.Context(), svcCtx)
		resp := l.List()
		httpx.OkJson(w, resp)
	}
}
