package apisix

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
)

type MyCustomClaims struct {
	UserId string `json:"key"`
	jwt.RegisteredClaims
}

func AuthMiddleware(context *Init.Data) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := strings.Join(r.Header["Authorization"], "")
			prefix := "Bearer "
			tokenString := ""

			if auth != "" && strings.HasPrefix(auth, prefix) {
				tokenString = auth[len(prefix):]
			}

			if tokenString == "" {
				resp := new(data.Rest).Error("token不能为空！", nil)
				bs, _ := json.Marshal(resp)
				w.Write(bs)
				return
			}

			token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(""), nil
			})

			var userId string
			if claims, ok := token.Claims.(*MyCustomClaims); ok {
				userId = claims.UserId
			}

			if userId == "" {
				resp := new(data.Rest).Error("您还没有登录，请先登录", nil)
				bs, _ := json.Marshal(resp)
				w.Write(bs)
				return
			}
			context.Set("userId", userId)
			next(w, r)
		}
	}
}
