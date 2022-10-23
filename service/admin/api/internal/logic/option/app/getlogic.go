package app

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
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
	options := model.SiteInfo{}
	json.Unmarshal([]byte(option.OptionValue), &options)

	resp.Success("获取成功", options)
	return
}
