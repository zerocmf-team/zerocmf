/**
** @创建时间: 2022/3/5 13:30
** @作者　　: return
** @描述　　:
 */

package assets

import (
	"context"
	"errors"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Assets struct {
	UnimplementedAssetsServer
}

func (g *Assets) GetFileUrl(filePath string) (data *Data, err error) {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}
	services, _, err := client.Health().Service("gincmf", "", true, nil)
	if err != nil {
		return
	}
	var target string

	if len(services) == 0 {
		err = errors.New("主服务已停止，请联系管理员修复")
		return
	}

	for k, v := range services {
		address := v.Service.Address
		meta := v.Service.Meta
		// 暂时处理
		if k == 0 {
			target = address + ":" + meta["grpc"]
		}
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	c := NewAssetsClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPrevPath(ctx, &AssetsRequest{
		FilePath: filePath,
	})

	if err != nil {
		return
	}

	data = r.GetData()
	return

}
