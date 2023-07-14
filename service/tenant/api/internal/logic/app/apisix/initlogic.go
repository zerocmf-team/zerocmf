package apisix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"zerocmf/common/bootstrap/util"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitLogic {
	return &InitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type nodes struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}

type timeout struct {
	Connect int `json:"connect"`
	Send    int `json:"send"`
	Read    int `json:"read"`
}

type keepAlivePool struct {
	IdleTimeout int `json:"idle_timeout"`
	Requests    int `json:"requests"`
	Size        int `json:"size"`
}

type upstream struct {
	Id            string        `json:"id"`
	Nodes         []nodes       `json:"nodes"`
	Timeout       timeout       `json:"timeout"`
	Type          string        `json:"type"`
	Scheme        string        `json:"scheme"`
	PassHost      string        `json:"pass_host"`
	Name          string        `json:"name"`
	KeepalivePool keepAlivePool `json:"keepalive_pool"`
}

func (l *InitLogic) Init(req *types.ApisixReq) (resp types.Response) {

	// 创建上游列表

	c := l.svcCtx
	key := c.Config.Apisix.ApiKey
	host := c.Config.Apisix.Host

	header := map[string]string{"X-API-KEY": key}

	up := []upstream{
		{
			Id: "admin-api",
			Nodes: []nodes{
				{
					Host:   "192.168.1.239",
					Port:   8800,
					Weight: 1,
				},
			},
			Timeout: timeout{
				Connect: 30,
				Send:    30,
				Read:    30,
			},
			Type:     "roundrobin",
			Scheme:   "http",
			PassHost: "pass",
			Name:     "admin-api",
			KeepalivePool: keepAlivePool{
				IdleTimeout: 60,
				Requests:    1000,
				Size:        320,
			},
		},
	}

	for _, v := range up {

		var body bytes.Buffer
		err := json.NewEncoder(&body).Encode(v)
		if err != nil {
			resp.Error("系统出错", err.Error())
			return
		}

		code, bytes := util.Request("PUT", "http://"+host+":9180/apisix/admin/upstreams/"+v.Id, &body, header)
		fmt.Println(code, string(bytes))
	}

	resp.Success("操作成功！", up)
	return
}
