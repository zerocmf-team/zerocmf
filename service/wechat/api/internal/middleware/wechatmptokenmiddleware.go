package middleware

import (
	goRedis "github.com/go-redis/redis"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	weData "zerocmf/service/wechat/api/data"
	"zerocmf/service/wechat/api/internal/types"
	weUtil "zerocmf/service/wechat/api/util"
)

type WechatMpTokenMiddleware struct {
	Redis *goRedis.Client
	weData.MpInfo
	*Init.Data
}

func NewWechatMpTokenMiddleware(data *Init.Data, Redis *goRedis.Client) *WechatMpTokenMiddleware {
	return &WechatMpTokenMiddleware{
		Data:  data,
		Redis: Redis,
	}
}

func (m *WechatMpTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		redis := m.Redis
		resp := new(types.Response)
		appId := "wxce4c356a74b76720"
		secret := "df5cca115cf3db90736952d51a51a4c7"

		m.AppId = appId
		m.Secret = secret

		token, err := weUtil.MpToken(redis, appId, secret, false)

		if err != nil {
			resp.Error("系统错误！请联系管理员或稍后重试", err.Error())
		}

		m.Set("token", token)

		next(w, r)
	}
}
