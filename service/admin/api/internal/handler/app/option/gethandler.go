package option

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/app/option"
	"zerocmf/service/admin/api/internal/svc"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := option.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get()
		httpx.OkJson(w, resp)
	}
}
