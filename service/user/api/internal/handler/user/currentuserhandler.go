package user

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user"
	"gincmf/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewCurrentUserLogic(r.Context(), svcCtx)
		resp, err := l.CurrentUser()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
