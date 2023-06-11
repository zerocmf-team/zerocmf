package oauth

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/user/api/internal/logic/oauth"
	"zerocmf/service/user/api/internal/svc"
)

func TokenRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := oauth.NewTokenRequestLogic(r.Context(), svcCtx)
		resp := l.TokenRequest()
		httpx.OkJson(w, resp)
	}
}
