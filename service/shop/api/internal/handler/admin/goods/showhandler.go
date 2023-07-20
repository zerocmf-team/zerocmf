package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/shop/api/internal/logic/admin/goods"
	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewShowLogic(r, svcCtx)
		resp := l.Show(&req)
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
