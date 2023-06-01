package middleware

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
)

func NewSiteMiddleware(c *Init.Data) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			siteIdForm := r.Form["siteId"]
			siteId := ""
			if len(siteIdForm) > 0 {
				siteId = siteIdForm[0]
			}
			if siteId == "" {
				resp := new(data.Rest).Error("站点不存在！", nil)
				bs, _ := json.Marshal(resp)
				w.Write(bs)
				return
			}
			c.Set("siteId", siteId)
			// 需要进一步验证站点信息是否存在

			next(w, r)
		}
	}
}
