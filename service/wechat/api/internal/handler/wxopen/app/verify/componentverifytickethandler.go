package verify

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/wxopen/app/verify"
	"zerocmf/service/wechat/api/internal/svc"
)

func ComponentVerifyTicketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := verify.NewComponentVerifyTicketLogic(r, svcCtx)
		resp := l.ComponentVerifyTicket()
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
