package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/redis"

	"github.com/jinzhu/copier"
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"github.com/zeromicro/go-zero/rest"
)

func ComponentAccessTokenMiddleware(context *Init.Data, redis redis.Redis) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			options := wxopen.GetOption()
			ticket := options.ComponentVerifyTicket
			var (
				accessToken string
				err         error
			)
			if ticket == "" {
				ticket, _ = redis.Get("componentVerifyTicket").Result()
			}

			if ticket != "" {
				accessToken, err = redis.Get("componentAccessToken").Result()
				if accessToken == "" || err != nil {
					bizContent := wxopen.ComponentAccessToken{}
					err = copier.Copy(&bizContent, &options)
					bizContent.ComponentVerifyTicket = ticket
					if err != nil {
						resp := new(data.Rest).ToBytes("系统错误", err.Error())
						w.Write(resp)
						return
					}
					var result wxopen.AccessToken
					result, err = new(wxopen.Component).ComponentAccessToken(bizContent)
					if err != nil {
						resp := new(data.Rest).ToBytes("系统错误", err.Error())
						w.Write(resp)
					}

					if result.ErrCode == 0 {
						accessToken = strings.TrimSpace(result.AccessToken)
						if accessToken != "" {
							redis.Set("componentAccessToken", accessToken, time.Minute*110)
						}
					}
				}
				if accessToken != "" {
					context.Set("componentAccessToken", accessToken)
				}
			}
			fmt.Println("accessToken", accessToken)
			next(w, r)
		}
	}
}
