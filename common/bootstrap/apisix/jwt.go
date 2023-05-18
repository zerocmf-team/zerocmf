package apisix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"zerocmf/common/bootstrap/util"
)

type (
	Jwt struct {
		apisix
	}
)

func NewJwt(host string) (j Jwt) {
	j = Jwt{apisix{Host: host}}
	return
}

//curl "http://127.0.0.1:9080/apisix/plugin/jwt/sign?key=1" -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X GET

func (r Jwt) GetAuthorizeToken(userId string) (token string, err error) {

	//bytes, err := json.Marshal(&payload)
	//if err != nil {
	//	fmt.Println("err", err)
	//}

	url := "http://" + r.Host + ":9080/apisix/plugin/jwt/sign?key=" + userId
	//if string(bytes) != "" {
	//	url += "&payload=" + string(bytes)
	//}

	code, resBytes := util.Request("GET", url, nil, nil)
	if !strings.HasPrefix(strconv.Itoa(code), "20") {
		err = errors.New("consumer is expired")
		fmt.Println("res", string(resBytes))
		return
	}
	token = string(resBytes)
	return
}
