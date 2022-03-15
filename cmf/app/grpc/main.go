/**
** @创建时间: 2021/11/24 22:44
** @作者　　: return
** @描述　　:
 */

package grpc

import (
	"gincmf/app/grpc/assets"
	"gincmf/app/grpc/user"
	"github.com/gincmf/bootstrap/consul"
	cmfGrpc "github.com/gincmf/bootstrap/grpc"
	"google.golang.org/grpc"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 注册grpc
 * @Date 2021/11/24 22:52:26
 * @Param
 * @return
 **/

func init() {
	registerConsul()
	cmfGrpc.NewServer(Register)
}

func registerConsul()  {
	consul.Register()
}

func Register(s *grpc.Server) {
	assets.RegisterAssetsServer(s, &assets.Assets{})
	user.RegisterUserServer(s, &user.User{})
}
