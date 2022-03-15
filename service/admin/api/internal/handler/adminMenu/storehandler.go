package adminMenu

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/adminMenu"
	"gincmf/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewStoreLogic(r.Context(), svcCtx)
		resp, err := l.Store()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
