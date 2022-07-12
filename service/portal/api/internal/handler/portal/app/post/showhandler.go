package post

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/app/post"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := post.NewShowLogic(r.Context(), svcCtx)
		resp := l.Show(&req)
		httpx.OkJson(w, resp)
	}
}
