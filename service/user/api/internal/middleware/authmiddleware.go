package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/data"
)

type AuthMiddleware struct {
	*Init.Data
}

func NewAuthMiddleware(data *Init.Data) *AuthMiddleware {
	return &AuthMiddleware{Data: data}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := apisix.Middleware{Data: m.Data}.Handle(w, r)
		if err != nil {
			new(data.Rest).Error("您还没有登录，请先登录", nil)
			httpx.OkJson(w, err)
		}
		next(w, r)
	}
}
