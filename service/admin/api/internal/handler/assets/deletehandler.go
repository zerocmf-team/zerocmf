package assets

import (
	"net/http"

	"zerocmf/service/admin/api/internal/logic/assets"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
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
