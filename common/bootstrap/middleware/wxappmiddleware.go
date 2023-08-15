package middleware

import (
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/tenant/rpc/tenantclient"
)

type WxappMiddleware struct {
	*Init.Data
	TenantRpc tenantclient.Tenant
}

func NewWxappMiddleware(data *Init.Data, tenantRpc tenantclient.Tenant) *WxappMiddleware {
	return &WxappMiddleware{
		Data:      data,
		TenantRpc: tenantRpc,
	}
}

// 根据appId获取siteId

func (m *WxappMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tenantRpc := m.TenantRpc
		r.ParseForm()
		appId := strings.Join(r.Form["appId"], "")
		showResult, err := tenantRpc.ShowMp(r.Context(), &tenantclient.ShowMpData{
			AppId: appId,
		})
		if err != nil {
			new(data.Rest).ToBytes("rpc服务错误！", err.Error())
			return
		}
		siteId := showResult.GetSiteId()
		m.Set("siteId", siteId)
		next(w, r)
	}
}
