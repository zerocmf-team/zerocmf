package list

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/app/list"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleSearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := list.NewSearchLogic(r.Context(), svcCtx)
		resp := l.Search(&req)
		httpx.OkJson(w, resp)
	}
}
