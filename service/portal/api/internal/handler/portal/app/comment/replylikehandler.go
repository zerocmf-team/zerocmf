package comment

import (
	"net/http"

	"zerocmf/service/portal/api/internal/logic/portal/app/comment"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewReplyLikeLogic(r.Context(), svcCtx)
		resp := l.ReplyLike(&req)
		httpx.OkJson(w, resp)
	}
}
