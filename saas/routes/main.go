/**
** @创建时间: 2021/11/23 13:10
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"gincmf/bootstrap/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func init() {

	conf := config.Config()
	addr := ":8080"
	if conf.App.Port != "" {
		addr = conf.App.Port
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func router() http.Handler {
	e := gin.Default()
	ApiListen(e)
	return e
}
