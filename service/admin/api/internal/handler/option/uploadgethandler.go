package option

import (
	"net/http"

	"gincmf/service/admin/api/internal/logic/option"
	"gincmf/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := option.NewUploadGetLogic(r.Context(), svcCtx)
		resp, _ := l.UploadGet()
		httpx.OkJson(w, resp)
	}
}
