package apisix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"zerocmf/common/bootstrap/util"
)

type Meta struct {
	Disable bool `json:"disable,omitempty"`
}

type JWTAuth struct {
	Meta Meta `json:"_meta,omitempty"`
}

type ProxyRewrite struct {
	RegexURI []string `json:"regex_uri,omitempty"`
}

type RoutePlugins struct {
	JWTAuth      *JWTAuth      `json:"jwt-auth,omitempty"`
	ProxyRewrite *ProxyRewrite `json:"proxy-rewrite,omitempty"`
}

type Route struct {
	URI       string       `json:"uri"`
	Name      string       `json:"name"`
	Methods   []string     `json:"methods,omitempty"`
	ServiceID string       `json:"service_id"`
	Plugins   RoutePlugins `json:"plugins"`
	Status    int          `json:"status"`
}

func (r *Route) Register(host string, apiKey string, routes []Route) (err error) {
	for _, v := range routes {
		var body bytes.Buffer
		err = json.NewEncoder(&body).Encode(&v)
		if err != nil {
			return
		}

		header := map[string]string{"X-API-KEY": apiKey}
		code, response := util.Request("PUT", "http://"+host+":9180/apisix/admin/routes/"+v.Name, &body, header)
		if !(code == 201 || code == 200) {
			err = errors.New(v.Name + ":" + string(response))
			return
		}
		fmt.Println("register routes "+v.Name+":", string(response))
	}
	return
}
