package route

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/app/route"
	"zerocmf/service/portal/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := route.NewListLogic(r, svcCtx)
		resp := l.List()
		httpx.OkJson(w, resp)
	}
}
