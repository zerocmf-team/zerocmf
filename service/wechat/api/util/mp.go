/**
** @创建时间: 2022/9/2 10:24
** @作者　　: return
** @描述　　:
 */

package util

import (
	goRedis "github.com/go-redis/redis"
	"github.com/zerocmf/wechat-easy-sdk/mp/base"
	"time"
)

const MpTokenKey = "mp_token"

func MpToken(redis *goRedis.Client, appId string, secret string, reload bool) (token string, err error) {
	token = redis.Get(MpTokenKey).Val()
	if reload {
		token = ""
	}
	if token == "" {
		var res base.TokenResponse
		res, err = base.Token(appId, secret)
		if err != nil {
			return
		}
		token = res.AccessToken
		expires := time.Second * time.Duration(res.ExpiresIn)
		redis.Set(MpTokenKey, token, expires)
	}
	return
}
