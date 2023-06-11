package account

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/user/api/internal/logic/admin/account"
	"zerocmf/service/user/api/internal/svc"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := account.NewCurrentUserLogic(r.Context(), svcCtx)
		resp := l.CurrentUser()
		httpx.OkJson(w, resp)
	}
}
