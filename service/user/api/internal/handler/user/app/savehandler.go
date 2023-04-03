package app

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/user/api/internal/logic/user/app"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
)

func SaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := app.NewSaveLogic(r.Context(), svcCtx)
		resp := l.Save(&req)
		httpx.OkJson(w, resp)
	}
}
