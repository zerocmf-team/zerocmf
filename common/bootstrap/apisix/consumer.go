package apisix

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/apisix/plugins/authentication"
	"zerocmf/common/bootstrap/util"
)

type (
	consumer struct {
		Apisix
	}
	pluginOption func(p *plugin)

	// 单个插件集合
	plugin struct {
		Name   string
		Config map[string]interface{}
	}

	Data struct {
		Username string                 `json:"username"`
		Plugins  map[string]interface{} `json:"plugins"`
	}
)

func NewConsumer(apiKey string, host string) (c consumer) {
	c = consumer{Apisix{ApiKey: apiKey, Host: host}}
	return
}

func pluginFunc(p *plugin, ins interface{}) {
	t := reflect.TypeOf(ins)
	v := reflect.ValueOf(ins)
	_type := reflect.TypeOf(ins)
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := _type.Field(i)
		tag := fieldType.Tag
		key := t.Field(i).Name
		if tag.Get("json") != "" {
			tagStr := tag.Get("json")
			tagArr := strings.Split(tagStr, ",")
			key = tagArr[0]
		}
		if field.IsValid() && !field.IsZero() {
			p.Config[key] = field.Interface()
		}
	}
	return
}

func WithJwtAuth(jwtPlugin authentication.JwtAuth) pluginOption {
	return func(p *plugin) {
		p.Name = "jwt-auth"
		pluginFunc(p, jwtPlugin)
	}
}

func (c consumer) Add(username string, opts ...pluginOption) (err error) {
	data := new(Data)
	data.Username = username
	data.Plugins = make(map[string]interface{}, 0)
	for _, option := range opts {
		p := plugin{Config: make(map[string]interface{}, 0)}
		option(&p)
		data.Plugins[p.Name] = p.Config
	}

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(data)
	if err != nil {
		return
	}

	header := map[string]string{"X-API-KEY": c.ApiKey}
	host := c.Host
	code, resBytes := util.Request("PUT", "http://"+host+":9180/apisix/admin/consumers", &body, header)
	if !strings.HasPrefix(strconv.Itoa(code), "20") {
		err = errors.New("errcode:" + strconv.Itoa(code) + string(resBytes))
		return
	}
	return
}
