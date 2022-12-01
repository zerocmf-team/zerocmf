package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/common/bootstrap/Init"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/config"
	"zerocmf/service/admin/api/internal/middleware"
	"zerocmf/service/admin/model"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	Request *http.Request
	*Init.Data
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	database := database.NewDb(c.Database)
	// 数据库迁移
	curDb := database.Db()
	model.Migrate("", true)

	return &ServiceContext{
		Config:         c,
		Db:             curDb,
		Data:           new(Init.Data).Context(),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
