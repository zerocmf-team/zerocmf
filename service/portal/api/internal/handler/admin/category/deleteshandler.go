package category

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/category"
	"zerocmf/service/portal/api/internal/svc"
)

func DeletesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := category.NewDeletesLogic(r, svcCtx)
		resp := l.Deletes()
		httpx.OkJson(w, resp)
	}
}
