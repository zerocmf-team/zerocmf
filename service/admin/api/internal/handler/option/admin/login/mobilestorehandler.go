package login

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/common/bootstrap/data"
	cmfValidator "zerocmf/common/bootstrap/validator"
	"zerocmf/service/admin/api/internal/logic/option/admin/login"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
)

func MobileStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MobileLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		zhValidator := new(cmfValidator.Zh).Validator()
		ZhTrans := new(cmfValidator.Zh).Trans()
		rest := new(data.Rest)
		err := zhValidator.Struct(req)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			for _, e := range errs {
				msg := rest.Error(e.Translate(ZhTrans), nil)
				httpx.OkJson(w, msg)
				return
			}
			msg := rest.Error(err.Error(), nil)
			httpx.OkJson(w, msg)
			return
		}

		l := login.NewMobileStoreLogic(r.Context(), svcCtx)
		resp := l.MobileStore(&req)
		httpx.OkJson(w, resp)
	}
}
