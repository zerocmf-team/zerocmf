package navItem

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/service/portal/api/internal/logic/admin/navItem"
	"zerocmf/service/portal/api/internal/svc"
)

func OptionsUrlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := navItem.NewOptionsUrlsLogic(r, svcCtx)
		resp := l.OptionsUrls()
		httpx.OkJson(w, resp)
	}
}
