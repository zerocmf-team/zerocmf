package mp

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/mp"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
)

func CheckSignatureHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckSignatureReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mp.NewCheckSignatureLogic(r.Context(), svcCtx)
		resp := l.CheckSignature(&req)
		if resp.Code == 1 {
			switch resp.Data.(type) {
			case string:
				echoStr := resp.Data.(string)
				w.Write([]byte(echoStr))
			default:
				w.Write([]byte("类型错误"))
			}
			httpx.Ok(w)
			return
		}
		httpx.OkJson(w, resp)
	}
}
