/**
** @创建时间: 2021/11/24 22:50
** @作者　　: return
** @描述　　:
 */

package grpc

import (
	"github.com/gincmf/bootstrap/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	register func(s *grpc.Server)
)

func NewServer(handle func(s *grpc.Server)) {
	register = handle
}

func ListenAndServe() error {
	conf := config.Config()
	lis, err := net.Listen("tcp", ":"+conf.Grpc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	if register != nil {
		register(s)
	}
	log.Printf("grpc server listening at %v", lis.Addr())
	return s.Serve(lis)
}
