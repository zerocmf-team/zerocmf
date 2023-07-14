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
	UserId string `json:"userId"`
	Type   string `json:"type"`
	SiteId string `json:"siteId"`
	Key    string `json:"key"`
	jwt.RegisteredClaims
}

func AuthMiddleware(context *Init.Data, tenantRpc tenantclient.Tenant) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
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
			var loginType string
			var userSiteId string

			if claims, ok := token.Claims.(*MyCustomClaims); ok {
				userId = claims.UserId
				loginType = claims.Type
				userSiteId = claims.SiteId

				if userId == "" {
					userId = claims.Key
				}
			}

			if userId == "" {
				resp := new(data.Rest).Error("您还没有登录，请先登录", nil)
				bs, _ := json.Marshal(resp)
				w.Write(bs)
				return
			}

			if loginType == "ram" {

				siteId := r.URL.Query().Get("siteId")
				if userSiteId != siteId {
					resp := new(data.Rest).Error("您没有该站点的访问权限！", nil)
					bs, _ := json.Marshal(resp)
					w.Write(bs)
					return
				}

				context.Set("userId", userId)
				context.Set("loginType", loginType)
				context.Set("siteId", userSiteId)

			} else {
				iSiteId, exist := context.Get("siteId")
				context.Set("userId", userId)
				if exist && tenantRpc != nil {
					//根据uid获取当前oid
					tenantReply, err := tenantRpc.Get(r.Context(), &tenant.CurrentUserReq{Uid: userId, SiteId: iSiteId.(string)})
					if err != nil {
						resp := new(data.Rest).Error("系统错误", err.Error())
						bs, _ := json.Marshal(resp)
						w.Write(bs)
						return
					}
					userId = strconv.FormatInt(tenantReply.Oid, 10)
					context.Set("userId", userId)
				}
			}
			next(w, r)
		}
	}
}
