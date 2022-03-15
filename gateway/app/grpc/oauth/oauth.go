/**
** @创建时间: 2021/12/17 19:10
** @作者　　: return
** @描述　　:
 */

package oauth

import (
	"context"
	"errors"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Oauth struct {
	UnimplementedOauthServer
}

func (s *Oauth) Request(token string) (userId string, err error) {

	// 服务发现查找
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}
	services,_,err := client.Health().Service("user","",true,nil)
	if err != nil {
		return
	}
	var target string

	if len(services) == 0 {
		err = errors.New("用户服务已停止，请联系管理员修复")
		return
	}

	for k,v := range services{
		address := v.Service.Address
		meta := v.Service.Meta
		// 暂时处理
		if k == 0 {
			target = address +":"+ meta["grpc"]
		}
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	c := NewOauthClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ValidationBearerToken(ctx, &OauthRequest{Token: token})
	if err != nil {
		return
	}
	//code := r.GetCode()
	//if code == 0 {
	//	err = errors.New(r.GetMessage())
	//	return
	//}
	userId = r.GetUserId()
	return
}
