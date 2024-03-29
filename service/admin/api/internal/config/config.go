package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zerocmf/common/bootstrap/apisix"
	"zerocmf/common/bootstrap/database"
)

type Config struct {
	rest.RestConf
	Etcd      discov.EtcdConf `json:",optional,inherit"`
	UserRpc   zrpc.RpcClientConf
	TenantRpc zrpc.RpcClientConf
	App       struct {
		Domain string `json:",optional"`
	}
	Database database.Database
	Apisix   apisix.Apisix
}
