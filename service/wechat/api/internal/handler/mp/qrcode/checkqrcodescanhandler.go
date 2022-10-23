package qrcode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/wechat/api/internal/logic/mp/qrcode"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
)

func CheckQrcodeScanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckQrcodeScanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := qrcode.NewCheckQrcodeScanLogic(r.Context(), svcCtx)
		resp := l.CheckQrcodeScan(&req)
		httpx.OkJson(w, resp)
	}
}
