package breadcrumb

import (
	"net/http"

	"gincmf/service/portal/api/internal/logic/portal/app/breadcrumb"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BreadcrumbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := breadcrumb.NewBreadcrumbLogic(r.Context(), svcCtx)
		resp := l.Breadcrumb(&req)
		httpx.OkJson(w, resp)
	}
}
