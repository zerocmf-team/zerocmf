package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/shop/api/internal/logic/admin/category"
	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"
)

func StoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategorySaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewStoreLogic(r, svcCtx)
		resp := l.Store(&req)
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
