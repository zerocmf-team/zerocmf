package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/apisix"
)

type Config struct {
	rest.RestConf
	Etcd      discov.EtcdConf `json:",optional,inherit"`
	TenantRpc zrpc.RpcClientConf
	ShopRpc   zrpc.RpcClientConf
	Apisix    apisix.Apisix
}
