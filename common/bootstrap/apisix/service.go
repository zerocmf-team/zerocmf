package apisix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"zerocmf/common/bootstrap/util"
)

type Service struct {
	Name       string `json:"name"`
	UpstreamId string `json:"upstream_id"`
}

func (s *Service) Register(host string, apiKey string) (err error) {
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(&s)
	if err != nil {
		return
	}
	header := map[string]string{"X-API-KEY": apiKey}
	code, response := util.Request("PUT", "http://"+host+":9180/apisix/admin/services/"+s.Name, &body, header)
	if code == 201 || code == 200 {
		fmt.Println("register service "+s.Name+":", string(response))
		return
	}
	err = errors.New(string(response))
	return
}
