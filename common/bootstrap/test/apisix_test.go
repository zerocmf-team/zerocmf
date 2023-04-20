package test

import (
	"fmt"
	"testing"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/apisix/plugins/authentication"
)

func TestConsumer(t *testing.T) {
	err := apisix.NewConsumer("edd1c9f034335f136f87ad84b625c8f1").Add("aaa", apisix.WithJwtAuth(authentication.JwtAuth{Key: "aaa"}))
	if err != nil {
		fmt.Println("err", err)
	}
}

func TestJwt(t *testing.T) {
	token, err := apisix.NewJwt("edd1c9f034335f136f87ad84b625c8f1").GetAuthorizeToken("aaa")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("token", token)
}
