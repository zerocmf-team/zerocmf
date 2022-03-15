/**
** @创建时间: 2021/11/24 22:44
** @作者　　: return
** @描述　　:
 */

package grpc

import (
	"gincmf/app/grpc/oauth"
	cmfGrpc "github.com/gincmf/bootstrap/grpc"
	"google.golang.org/grpc"
)

func init() {
	cmfGrpc.NewServer(Register)
}

func Register(s *grpc.Server) {
	oauth.RegisterOauthServer(s, &oauth.Oauth{})
}
