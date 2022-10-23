package admin

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadGetLogic {
	return &UploadGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadGetLogic) UploadGet() (resp types.Response) {
	c := l.svcCtx
	db := c.Db

	option := model.Option{}
	uploadSetting := model.UploadSetting{}
	tx := db.Where("option_name = ?", "upload_setting").First(&option)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}
	value := option.OptionValue
	err := json.Unmarshal([]byte(value), &uploadSetting)

	if err != nil {
		resp.Error("解析时出错："+err.Error(), nil)
		return
	}

	resp.Success("获取成功！", uploadSetting)
	return
}
