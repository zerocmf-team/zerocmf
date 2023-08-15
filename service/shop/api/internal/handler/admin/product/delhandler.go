package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/shop/api/internal/logic/admin/product"
	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"
)

func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewDelLogic(r, svcCtx)
		resp := l.Del(&req)
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}

		if resp.OkBytes() {
			w.Write([]byte(resp.Msg))
			httpx.Ok(w)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
