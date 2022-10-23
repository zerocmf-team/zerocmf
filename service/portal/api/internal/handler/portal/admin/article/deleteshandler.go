package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/portal/api/internal/logic/portal/admin/article"
	"zerocmf/service/portal/api/internal/svc"
)

func DeletesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article.NewDeletesLogic(r.Context(), svcCtx)
		resp := l.Deletes()
		httpx.OkJson(w, resp)
	}
}
