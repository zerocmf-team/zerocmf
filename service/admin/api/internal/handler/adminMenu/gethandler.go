package adminMenu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/logic/adminMenu"
	"zerocmf/service/admin/api/internal/svc"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get()
		httpx.OkJson(w, resp)
	}
}
