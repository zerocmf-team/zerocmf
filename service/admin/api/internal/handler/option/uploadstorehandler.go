package option

import (
	"net/http"

	"zerocmf/service/admin/api/internal/logic/option"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := option.NewUploadStoreLogic(r.Context(), svcCtx)
		resp, _ := l.UploadStore(&req)
		httpx.OkJson(w, resp)
	}
}
