package qrcode

import (
	"net/http"
	"zerocmf/service/wechat/api/internal/logic/mp/qrcode"
	"zerocmf/service/wechat/api/internal/svc"
)

func WsQrcodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := qrcode.NewWsQrcodeLogic(r.Context(), svcCtx)
		l.WsQrcode(w,r)
	}
}
