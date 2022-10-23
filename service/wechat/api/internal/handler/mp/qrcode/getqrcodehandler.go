package qrcode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/mp/qrcode"
	"zerocmf/service/wechat/api/internal/svc"
)

func GetQrcodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := qrcode.NewGetQrcodeLogic(r.Context(), svcCtx)
		resp := l.GetQrcode()
		httpx.OkJson(w, resp)
	}
}
