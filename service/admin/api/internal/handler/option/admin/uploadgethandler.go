package admin

import (
	"net/http"
	"zerocmf/service/admin/api/internal/logic/admin/option/admin"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/admin/api/internal/svc"
)

func UploadGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewUploadGetLogic(r.Context(), svcCtx)
		resp := l.UploadGet()
		httpx.OkJson(w, resp)
	}
}
