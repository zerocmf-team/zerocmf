/**
** @创建时间: 2021/11/22 11:01
** @作者　　: return
** @描述　　:
 */
package bootstrap

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Rpc(callBack func(*grpc.Server)) {

	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	callBack(s)

	g.Go(func() error {
		log.Printf("server listening at %v", lis.Addr())
		err := s.Serve(lis)
		return err
	})

}
