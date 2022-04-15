package app

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/app"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := app.NewSaveLogic(r.Context(), svcCtx)
		resp, _ := l.Save(&req)
		httpx.OkJson(w, resp)
	}
}
