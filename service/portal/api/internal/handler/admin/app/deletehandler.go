package app

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/app"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := app.NewDeleteLogic(r, svcCtx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}
