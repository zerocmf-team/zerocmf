package account

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerocmf/service/user/api/internal/logic/user/admin/account"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		fmt.Println("req",req)
		l := account.NewGetLogic(r.Context(), svcCtx)
		resp := l.Get(&req)
		httpx.OkJson(w, resp)
	}
}
