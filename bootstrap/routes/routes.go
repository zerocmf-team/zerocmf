/**
** @创建时间: 2021/11/24 22:55
** @作者　　: return
** @描述　　:
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/config"
	"github.com/gincmf/bootstrap/grpc"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
	"strconv"
)

var (
	g      errgroup.Group
)

func NewRoutes(listen ...func(e *gin.Engine)) {

	conf := config.Config()
	addr := ":8080"
	if conf.App.Port > 0 {
		addr = ":" + strconv.Itoa(conf.App.Port)
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      router(listen...),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("http server listening at %v", addr)

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		return grpc.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func router(listens ...func(e *gin.Engine)) http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	if len(listens) > 0 {
		for _,listen := range listens {
			listen(e)
		}
	}
	return e
}
