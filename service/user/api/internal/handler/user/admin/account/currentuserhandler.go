package account

import (
	"net/http"
	"zerocmf/service/user/api/internal/logic/user/admin/account"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/user/api/internal/svc"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewCurrentUserLogic(r.Context(), svcCtx)
		resp := l.CurrentUser()
		httpx.OkJson(w, resp)
	}
}
