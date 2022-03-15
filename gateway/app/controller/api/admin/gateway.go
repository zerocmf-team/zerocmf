/**
** @创建时间: 2021/12/9 16:25
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Gateway struct {
	controller.Rest
}

func (rest *Gateway) Register(c *gin.Context) {

	requestURI := c.Request.RequestURI
	pathUrl := strings.Split(requestURI,"?")
	path := pathUrl[0]

	urlArr := strings.Split(requestURI, "/")
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	//转发的url，端口
	target := "127.0.0.1:4001"

	var moduleName string

	if len(urlArr) > 4 {
		//api v1 [app] [module]
		moduleName = urlArr[4]
	}

	reg := regexp.MustCompile(`/public`)
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	switch path {
	case "/api/oauth/token":
		moduleName = "user"
	case "/api/oauth/refresh":
		moduleName = "user"
	}

	if moduleName != "" {
		agent := client.Agent()
		health := client.Health()
		info, err := agent.Self()
		if err != nil {
			rest.Error(c, "consul发送错误："+err.Error(), nil)
			return
		}

		name := info["Config"]["NodeName"].(string)
		checks, _, err := health.Node(name, nil)

		for _, v := range checks {
			if v.ServiceID != "" {
				if moduleName == v.ServiceName {
					if v.Status != "passing" {
						rest.Error(c, v.ServiceName+"服务已经关闭，请联系管理员修复!", nil)
						return
					}
					service, _, err := agent.Service(v.ServiceID, nil)
					if err != nil {
						rest.Error(c, v.ServiceName+"服务出错，请联系管理员修复!", nil)
						return
					}
					port := strconv.Itoa(service.Port)
					target = service.Address + ":" + port
				}
			}
		}
	}

	userId, exist := c.Get("userId")

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = target
	query := u.Query()
	if exist {
		query.Add("userId", userId.(string))
	}
	u.RawQuery = query.Encode()
	proxy := httputil.NewSingleHostReverseProxy(u)

	//重写出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("http: proxy error: %v", err)
		ret := fmt.Sprintf("http proxy error %v", err)

		//写到body里
		rw.Write([]byte(ret))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
