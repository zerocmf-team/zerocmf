package adminMenu

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/admin/adminMenu"
	"zerocmf/service/admin/api/internal/svc"
)

func SyncHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminMenu.NewSyncLogic(r.Context(), svcCtx)
		resp := l.Sync()
		httpx.OkJson(w, resp)
	}
}
