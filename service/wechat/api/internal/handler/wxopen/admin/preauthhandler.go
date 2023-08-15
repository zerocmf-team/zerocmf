package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/wxopen/admin"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
)

func PreAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PreAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewPreAuthLogic(r, svcCtx)
		resp := l.PreAuth(&req)
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}

		if resp.OkBytes() {
			w.Write([]byte(resp.Msg))
			httpx.Ok(w)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
