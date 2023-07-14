package apisix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"zerocmf/common/bootstrap/util"
)

type nodes struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}

type timeout struct {
	Connect int `json:"connect"`
	Send    int `json:"send"`
	Read    int `json:"read"`
}

type keepAlivePool struct {
	IdleTimeout int `json:"idle_timeout"`
	Requests    int `json:"requests"`
	Size        int `json:"size"`
}

type Upstream struct {
	Nodes         []nodes       `json:"nodes"`
	Timeout       timeout       `json:"timeout"`
	Type          string        `json:"type"`
	Scheme        string        `json:"scheme"`
	PassHost      string        `json:"pass_host"`
	Name          string        `json:"name,optional"`
	KeepalivePool keepAlivePool `json:"keepalive_pool"`
}

/**
Desc: 注册上游服务
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-14 07:56:01
*/

func (u *Upstream) Register(host string, apiKey string) (err error) {
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(&u)
	if err != nil {
		return
	}
	header := map[string]string{"X-API-KEY": apiKey}
	code, response := util.Request("PUT", "http://"+host+":9180/apisix/admin/upstreams/"+u.Name, &body, header)
	if code == 201 || code == 200 {
		fmt.Println("register upstream "+u.Name+":", string(response))
		return
	}
	err = errors.New(string(response))
	return
}
