package test

import (
	"fmt"
	"testing"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/apisix/plugins/authentication"
)

func TestConsumer(t *testing.T) {
	err := apisix.NewConsumer("edd1c9f034335f136f87ad84b625c8f1", "localhost").Add("aaa", apisix.WithJwtAuth(authentication.JwtAuth{Key: "aaa"}))
	if err != nil {
		fmt.Println("err", err)
	}
}

func TestJwt(t *testing.T) {
	//payload := map[string]interface{}{
	//	"tenantId": "1111",
	//	"oid":      "1",
	//}

	token, err := apisix.NewJwt("localhost").GetAuthorizeToken("1")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("token", token)
}
