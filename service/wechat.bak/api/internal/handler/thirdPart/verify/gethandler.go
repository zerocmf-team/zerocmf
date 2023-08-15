package verify

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/thirdPart/verify"
	"zerocmf/service/wechat/api/internal/svc"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := verify.NewGetLogic(r, svcCtx)
		resp := l.Get()
		if resp.StatusCode != nil {
			w.WriteHeader(*resp.StatusCode)
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
