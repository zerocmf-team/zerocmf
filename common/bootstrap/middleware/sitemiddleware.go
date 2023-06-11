package middleware

import (
	"encoding/json"
	"net/http"
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
		r.ParseForm()
		siteId := strings.Join(r.Form["siteId"], "")
		if siteId == "" {
			resp := new(data.Rest).Error("站点不存在！", nil)
			bs, _ := json.Marshal(resp)
			w.Write(bs)
			return
		}
		m.Set("siteId", siteId)
		next(w, r)
	}
}
