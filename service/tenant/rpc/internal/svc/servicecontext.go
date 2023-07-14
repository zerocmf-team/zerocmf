package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/rpc/adminclient"
	"zerocmf/service/tenant/rpc/internal/config"
)

type ServiceContext struct {
	Db       *gorm.DB
	Config   config.Config
	AdminRpc adminclient.Admin
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := database.NewGormDb(c.Database)
	return &ServiceContext{
		Config:   c,
		Db:       db,
		AdminRpc: adminclient.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
	}
}
