package themeFile

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/app/themeFile"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThemeFileListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := themeFile.NewListLogic(r.Context(), svcCtx)
		resp := l.List(&req)
		httpx.OkJson(w, resp)
	}
}
