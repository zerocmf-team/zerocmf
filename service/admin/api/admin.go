package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/router"
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/service/admin/api/internal/config"
	"zerocmf/service/admin/api/internal/handler"
	"zerocmf/service/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	r := router.NewRouter()
	r.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 资源映射
		if strings.HasPrefix(r.URL.Path, "/public") {
			handler := http.StripPrefix("/public",http.FileServer(http.Dir("public")))
			handler.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	}))
	server := rest.MustNewServer(c.RestConf,rest.WithRouter(r))
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
			Init.SetDomain(domain)
			ctx.Request = r
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
