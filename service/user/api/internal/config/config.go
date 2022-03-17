package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Database struct {
		Type     string
		Host     string
		Database string
		Username string
		Password string
		Port     int
		Charset  string
		Prefix   string
		AuthCode string
	}
}
