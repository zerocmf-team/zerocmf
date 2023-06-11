package apisix

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/tenant/rpc/types/tenant"
)

type MyCustomClaims struct {
	UserId string `json:"key"`
	jwt.RegisteredClaims
}

func AuthMiddleware(context *Init.Data, tenantRpc tenantclient.Tenant) rest.Middleware {
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
			siteId, exist := context.Get("siteId")
			context.Set("userId", userId)
			if exist && tenantRpc != nil {
				//根据uid获取当前oid
				tenantReply, err := tenantRpc.Get(r.Context(), &tenant.CurrentUserReq{Uid: userId, SiteId: siteId.(string)})
				if err != nil {
					resp := new(data.Rest).Error("系统错误", err.Error())
					bs, _ := json.Marshal(resp)
					w.Write(bs)
					return
				}
				userId = strconv.FormatInt(tenantReply.Oid, 10)
				context.Set("userId", userId)
			}
			next(w, r)
		}
	}
}
