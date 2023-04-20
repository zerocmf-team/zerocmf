package admin

import (
	"net/http"
	"zerocmf/service/admin/api/internal/logic/option/admin"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
)

func UploadStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewUploadStoreLogic(r.Context(), svcCtx)
		resp := l.UploadStore(&req)
		httpx.OkJson(w, resp)
	}
}
