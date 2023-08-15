package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
)

type SiteMiddleware struct {
	*Init.Data
}

func NewSiteMiddleware(data *Init.Data) *SiteMiddleware {
	return &SiteMiddleware{
		Data: data,
	}
}

func (m *SiteMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//小程序平台可能先拿到siteId，否则是h5。必须传入siteId
		_, exist := m.Get("siteId")
		if !exist {
			r.ParseForm()
			siteId := strings.Join(r.Form["siteId"], "")
			if siteId == "" {
				resp := new(data.Rest).Error("站点不存在！", nil)
				bs, _ := json.Marshal(resp)
				w.Write(bs)
				return
			}
			siteIdInt, err := strconv.ParseInt(siteId, 10, 64)
			if err != nil {
				new(data.Rest).ToBytes("非法站点", nil)
				return
			}
			m.Set("siteId", siteIdInt)
		}
		next(w, r)
	}
}
