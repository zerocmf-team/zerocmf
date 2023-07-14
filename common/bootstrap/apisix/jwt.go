package apisix

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/util"
)

type (
	Jwt struct {
		Apisix
	}
)

func NewJwt(host string) (j Jwt) {
	j = Jwt{Apisix{Host: host}}
	return
}

//curl "http://127.0.0.1:9080/apisix/plugin/jwt/sign?key=1" -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X GET

func (r Jwt) GetAuthorizeToken(userId string, payload map[string]string) (token string, err error) {
	queryParams := url.Values{}
	queryParams.Add("key", userId)
	var marshal []byte
	if payload != nil {
		marshal, err = json.Marshal(payload)
		payloadStr := string(marshal)
		queryParams.Add("payload", payloadStr)
	}
	url := "http://" + r.Host + ":9080/apisix/plugin/jwt/sign" + "?" + queryParams.Encode()

	fmt.Println("url", url)

	code, resBytes := util.Request("GET", url, nil, nil)
	if !strings.HasPrefix(strconv.Itoa(code), "20") {
		err = errors.New("consumer is expired")
		fmt.Println("res", url, string(resBytes))
		return
	}
	token = string(resBytes)
	return
}
