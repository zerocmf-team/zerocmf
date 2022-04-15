package route

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/route"
	"gincmf/service/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := route.NewListLogic(r.Context(), svcCtx)
		resp := l.List()
		httpx.OkJson(w, resp)
	}
}
