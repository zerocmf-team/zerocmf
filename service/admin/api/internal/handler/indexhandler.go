package handler

import (
	"net/http"

	"zerocmf/service/admin/api/internal/logic"
	"zerocmf/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp, _ := l.Index()
		httpx.OkJson(w, resp)
	}
}
