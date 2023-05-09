package adminMenu

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/admin/api/internal/logic/admin/adminMenu"
	"zerocmf/service/admin/api/internal/svc"
)

func GetAllMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		l := adminMenu.NewGetAllMenusLogic(r.Context(), svcCtx)
		resp := l.GetAllMenus()
		httpx.OkJson(w, resp)
	}
}
