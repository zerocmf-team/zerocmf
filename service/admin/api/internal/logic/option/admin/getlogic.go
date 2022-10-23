package admin

import (
	"context"
	"gorm.io/gorm"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	option := model.Option{}
	tx := db.Where("option_name = ?", "site_info").First(&option) // 查询
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		resp.Error("获取失败："+tx.Error.Error(), nil)
		return
	}
	resp.Success("获取成功", option)
	return
}
