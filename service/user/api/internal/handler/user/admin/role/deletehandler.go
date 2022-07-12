package role

import (
	"net/http"

	"zerocmf/service/user/api/internal/logic/user/admin/role"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleDelete
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewDeleteLogic(r.Context(), svcCtx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}
