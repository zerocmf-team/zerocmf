package category

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/category"
	"zerocmf/service/portal/api/internal/svc"
)

func OptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := category.NewOptionsLogic(r, svcCtx)
		resp := l.Options()
		httpx.OkJson(w, resp)
	}
}
