package assets

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/assets"
	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := assets.NewDeleteLogic(r.Context(), svcCtx)
		resp, _ := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}
