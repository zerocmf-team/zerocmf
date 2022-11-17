package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/logic"
	"zerocmf/service/admin/api/internal/svc"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp := l.Index()
		httpx.OkJson(w, resp)
	}
}
