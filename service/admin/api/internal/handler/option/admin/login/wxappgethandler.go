package login

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/option/admin/login"
	"zerocmf/service/admin/api/internal/svc"
)

func WxappGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewWxappGetLogic(r.Context(), svcCtx)
		resp := l.WxappGet()
		httpx.OkJson(w, resp)
	}
}
