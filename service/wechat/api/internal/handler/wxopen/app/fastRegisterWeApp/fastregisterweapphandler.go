package fastRegisterWeApp

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/wxopen/app/fastRegisterWeApp"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
)

func FastRegisterWeAppHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FastRegisterWeAppReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := fastRegisterWeApp.NewFastRegisterWeAppLogic(r, svcCtx)
		resp := l.FastRegisterWeApp(&req)
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
