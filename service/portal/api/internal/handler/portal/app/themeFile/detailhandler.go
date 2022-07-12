package themeFile

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/app/themeFile"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThemeFileDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := themeFile.NewDetailLogic(r.Context(), svcCtx)
		resp := l.Detail(&req)
		httpx.OkJson(w, resp)
	}
}
