package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/shop/api/internal/logic"
	"zerocmf/service/shop/api/internal/svc"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r, svcCtx)
		resp := l.Index()
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
