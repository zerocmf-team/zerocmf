/**
** @创建时间: 2022/2/26 09:43
** @作者　　: return
** @描述　　:
 */

package consul

import (
	"fmt"
	"github.com/gincmf/bootstrap/config"
	"github.com/hashicorp/consul/api"
	"strconv"
)

func Register() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}
	agent := client.Agent()

	conf := config.Config()
	app := conf.App
	grpc := conf.Grpc
	reg := &api.AgentServiceRegistration{
		Name: app.Name,
		Tags: []string{"local"},
		TaggedAddresses: map[string]api.ServiceAddress{
			"lan": {
				Address: "",
				Port:    80,
			},
		},
		Meta: map[string]string{
			"grpc": grpc.Port,
		},
		Port: app.Port,
		Check: &api.AgentServiceCheck{
			Name:          app.Name,
			HTTP:          "http://" + app.Domain + ":" + strconv.Itoa(app.Port) + "/ping",
			TLSSkipVerify: false,
			Method:        "GET",
			Interval:      "10s",
			Timeout:       "30s",
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Println("err", err.Error())
	}
}
