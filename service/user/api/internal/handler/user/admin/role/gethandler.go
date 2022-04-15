package role

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user/admin/role"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleGet
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewGetLogic(r.Context(), svcCtx)
		resp, _ := l.Get(&req)
		httpx.OkJson(w, resp)
	}
}
