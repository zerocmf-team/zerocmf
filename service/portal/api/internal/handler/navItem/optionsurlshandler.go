package navItem

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/portal/api/internal/logic/navItem"
	"zerocmf/service/portal/api/internal/svc"
)

func OptionsUrlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := navItem.NewOptionsUrlsLogic(r.Context(), svcCtx)
		resp := l.OptionsUrls()
		httpx.OkJson(w, resp)
	}
}
