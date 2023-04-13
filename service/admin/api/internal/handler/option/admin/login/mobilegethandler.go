package login

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/admin/option/admin/login"
	"zerocmf/service/admin/api/internal/svc"
)

func MobileGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewMobileGetLogic(r.Context(), svcCtx)
		resp := l.MobileGet()
		httpx.OkJson(w, resp)
	}
}
