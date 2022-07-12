package middleware

import (
	"zerocmf/common/bootstrap/data"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		r.ParseForm()
		userId := strings.Join(r.Form["userId"], "")

		if userId == "" {
			err := new(data.Rest).Error("您还没有登录，请先登录", nil)
			httpx.OkJson(w, err)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
