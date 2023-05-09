package article

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/portal/admin/article"
	"zerocmf/service/portal/api/internal/svc"
)

func DeletesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := article.NewDeletesLogic(r.Context(), svcCtx)
		resp := l.Deletes()
		httpx.OkJson(w, resp)
	}
}
