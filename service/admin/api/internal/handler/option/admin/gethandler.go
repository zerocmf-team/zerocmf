package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/logic/option/admin"
	"zerocmf/service/admin/api/internal/svc"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get()
		httpx.OkJson(w, resp)
	}
}
