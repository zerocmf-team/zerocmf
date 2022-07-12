package adminMenu

import (
	"net/http"

	"zerocmf/service/admin/api/internal/logic/adminMenu"
	"zerocmf/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
