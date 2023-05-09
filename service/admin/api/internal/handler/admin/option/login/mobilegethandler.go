package login

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/admin/option/login"
	"zerocmf/service/admin/api/internal/svc"
)

func MobileGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 获取请求头域名
		scheme := "http://"
		if r.Header.Get("Scheme") == "https" {
			scheme = "https://"
		}
		host := r.Host
		domain := scheme + host
		svcCtx.Config.App.Domain = domain
		svcCtx.Request = r
		l := login.NewMobileGetLogic(r.Context(), svcCtx)
		resp := l.MobileGet()
		httpx.OkJson(w, resp)
	}
}
