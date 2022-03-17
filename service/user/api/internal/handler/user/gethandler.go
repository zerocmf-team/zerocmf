package user

import (
	"net/http"

	"gincmf/service/user/api/internal/logic/user"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
