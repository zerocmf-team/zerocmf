package main

import (
	"flag"
	"fmt"
	"gincmf/common/bootstrap/data"
	"net/http"
	"strings"

	"gincmf/service/user/api/internal/config"
	"gincmf/service/user/api/internal/handler"
	"gincmf/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	data.SetSalts(c.Database.Prefix)

	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 获取请求头域名
			scheme := "http://"
			if r.Header.Get("Scheme") == "https" {
				scheme = "https://"
			}
			host := r.Host
			domain := scheme + host
			ctx.Config.App.Domain = domain
			data.SetDomain(domain)
			ctx.Request = r
			ctx.ResponseWriter = w
			// 处理userId
			r.ParseForm()
			userId := strings.Join(r.Form["userId"], "")
			ctx.Set("userId", userId)

			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
