package themeFile

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/admin/themeFile"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThemeFileSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := themeFile.NewSaveLogic(r.Context(), svcCtx)
		resp := l.Save(&req)
		httpx.OkJson(w, resp)
	}
}
