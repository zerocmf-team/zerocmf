package adminMenu

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/adminMenu"
	"gincmf/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewSyncLogic(r.Context(), svcCtx)
		resp, _ := l.Sync()
		httpx.OkJson(w, resp)
	}
}
