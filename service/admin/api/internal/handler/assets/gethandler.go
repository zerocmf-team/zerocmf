package assets

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/logic/assets"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := assets.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get(&req)
		httpx.OkJson(w, resp)
	}
}
