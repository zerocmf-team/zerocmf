package assets

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/assets"
	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := assets.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
