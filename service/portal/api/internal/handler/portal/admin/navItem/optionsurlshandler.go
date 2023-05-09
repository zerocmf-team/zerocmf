package navItem

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/portal/admin/navItem"
	"zerocmf/service/portal/api/internal/svc"
)

func OptionsUrlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := navItem.NewOptionsUrlsLogic(r.Context(), svcCtx)
		resp := l.OptionsUrls()
		httpx.OkJson(w, resp)
	}
}