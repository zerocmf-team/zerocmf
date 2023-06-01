package form

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/common/bootstrap/data"
	cmfValidator "zerocmf/common/bootstrap/validator"
	"zerocmf/service/lowcode/api/internal/logic/admin/form"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FormGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		ZhTrans, zhValidator := new(cmfValidator.Zh).Validator()
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

		// 获取请求头域名
		scheme := "http://"
		if r.Header.Get("Scheme") == "https" {
			scheme = "https://"
		}
		host := r.Host
		domain := scheme + host
		svcCtx.Config.App.Domain = domain
		svcCtx.Request = r
		l := form.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get(&req)
		httpx.OkJson(w, resp)
	}
}
