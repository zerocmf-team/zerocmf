package navItem

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/navItem"
	"zerocmf/service/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OptionsUrlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := navItem.NewOptionsUrlsLogic(r.Context(), svcCtx)
		resp := l.OptionsUrls()
		httpx.OkJson(w, resp)
	}
}
