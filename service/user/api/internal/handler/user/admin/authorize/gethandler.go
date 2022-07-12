package authorize

import (
	"net/http"

	"zerocmf/service/user/api/internal/logic/user/admin/authorize"
	"zerocmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := authorize.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get()
		httpx.OkJson(w, resp)
	}
}
