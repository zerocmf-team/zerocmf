package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestName(t *testing.T) {
	e, err := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")

	if err != nil {
		panic(err.Error())
	}

	sub := "2" // 想要访问资源的用户
	obj := "settings/index:api/settings/test" // 将要被访问的资源
	act := "get" // 用户对资源实施的操作

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理错误
		fmt.Println("err",err.Error())
	}

	fmt.Println("ok", ok)

	if ok == true {
		// 允许 alice 读取 data1
	} else {
		// 拒绝请求，抛出异常
	}


}