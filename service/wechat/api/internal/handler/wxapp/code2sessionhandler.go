package wxapp

import (
	"net/http"

	"zerocmf/service/wechat/api/internal/logic/wxapp"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Code2SessionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Code2SessionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := wxapp.NewCode2SessionLogic(r.Context(), svcCtx)
		resp := l.Code2Session(&req)
		httpx.OkJson(w, resp)
	}
}
