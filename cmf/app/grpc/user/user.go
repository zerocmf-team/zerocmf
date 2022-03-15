/**
** @创建时间: 2022/2/24 21:32
** @作者　　: return
** @描述　　:
 */

package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type User struct {
	UnimplementedUserServer
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取请求
 * @Date 2022/3/5 14:4:10
 * @Param
 * @return
 **/

func (s *User) Request(userId int, tenantId int) (data *Data, err error) {


	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}

	as,_,err := client.Agent().Service("user",nil)
	if err != nil {
		return
	}

	conn, err := grpc.Dial(":"+as.Meta["grpc"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &UserRequest{
		UserId: int32(userId),
		Tenant: int64(tenantId),
	})

	if err != nil {
		fmt.Println("err",err.Error())
		return
	}

	code := r.GetCode()
	msg := r.GetMessage()

	if code != 1 {
		err = errors.New(msg)
		return
	}

	data = r.GetData()
	return
}
