package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	App struct{
		Domain string `json:",optional"`
	}
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
