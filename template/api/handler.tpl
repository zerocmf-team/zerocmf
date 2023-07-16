package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r, svcCtx)
		resp := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if resp.StatusCode != nil {
            w.WriteHeader(*resp.StatusCode)
        }
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
