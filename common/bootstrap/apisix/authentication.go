package apisix

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/Init"
)

type MyCustomClaims struct {
	UserId string `json:"key"`
	jwt.RegisteredClaims
}

func AuthMiddleware(data *Init.Data) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := strings.Join(r.Header["Authorization"], "")
			prefix := "Bearer "
			tokenString := ""

			if auth != "" && strings.HasPrefix(auth, prefix) {
				tokenString = auth[len(prefix):]
			}

			if tokenString == "" {
				errors.New("token不能为空！")
				return
			}

			token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(""), nil
			})

			var userId string
			if claims, ok := token.Claims.(*MyCustomClaims); ok {
				fmt.Println("token.Valid", token.Valid)
				userId = claims.UserId
			}

			if userId == "" {
				errors.New("您还没有登录，请先登录")
				return
			}

			data.Set("userId", userId)
			next(w, r)
		}
	}
}
