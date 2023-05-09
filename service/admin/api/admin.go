package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/router"
	http "net/http"
	"strings"
	"zerocmf/common/bootstrap/middleware"
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
			httpHandler := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
			httpHandler.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	}))
	server := rest.MustNewServer(c.RestConf, rest.WithRouter(r))
	defer server.Stop()

	// 全局中间件
	server.Use(middleware.NewSiteMiddleware(ctx.Data).Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
