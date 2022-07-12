package config

import (
	"zerocmf/common/bootstrap/database"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Database database.Database
}
