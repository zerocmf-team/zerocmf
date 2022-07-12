package category

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/admin/category"
	"zerocmf/service/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewOptionsLogic(r.Context(), svcCtx)
		resp := l.Options()
		httpx.OkJson(w, resp)
	}
}
