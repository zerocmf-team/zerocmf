package assets

import (
	"net/http"
	"zerocmf/service/admin/api/internal/logic/assets"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := assets.NewDeleteLogic(r.Context(), svcCtx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}
