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

func NewJwt(apiKey string) (j Jwt) {
	j = Jwt{apisix{ApiKey: apiKey}}
	return
}

func (r Jwt) GetAuthorizeToken(userId string) (token string, err error) {
	code, resBytes := util.Request("GET", "http://localhost:9080/apisix/plugin/jwt/sign?key="+userId, nil, nil)
	if !strings.HasPrefix(strconv.Itoa(code), "20") {
		err = errors.New("consumer is expired")
		fmt.Println("res", string(resBytes))
		return
	}
	token = string(resBytes)
	return
}
