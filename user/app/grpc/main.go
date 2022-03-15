/**
** @创建时间: 2021/11/24 22:44
** @作者　　: return
** @描述　　:
 */

package grpc

import (
	"gincmf/app/grpc/assets"
	"gincmf/app/grpc/oauth"
	"gincmf/app/grpc/user"
	"github.com/gincmf/bootstrap/consul"
	cmfGrpc "github.com/gincmf/bootstrap/grpc"
	"google.golang.org/grpc"
)

func init() {
	cmfGrpc.NewServer(Register)
	registerConsul()
}

func registerConsul()  {
	consul.Register()
}

func Register(s *grpc.Server) {
	oauth.RegisterOauthServer(s, &oauth.Oauth{})
	user.RegisterUserServer(s, &user.User{})
	assets.RegisterAssetsServer(s, &assets.Assets{})
}
