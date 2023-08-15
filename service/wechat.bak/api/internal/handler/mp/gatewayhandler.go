package mp

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/mp"
	"zerocmf/service/wechat/api/internal/svc"
)

func GatewayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := mp.NewGatewayLogic(r.Context(), svcCtx)
		resp := l.Gateway()
		if resp.Code == 1 {
			switch resp.Data.(type) {
			case []byte:
				echoBytes := resp.Data.([]byte)
				w.Write(echoBytes)
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
