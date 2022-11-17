package adminMenu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/logic/adminMenu"
	"zerocmf/service/admin/api/internal/svc"
)

func SyncHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewSyncLogic(r.Context(), svcCtx)
		resp := l.Sync()
		httpx.OkJson(w, resp)
	}
}
