package login

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxappGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxappGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxappGetLogic {
	return &WxappGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxappGetLogic) WxappGet() (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	option := model.Option{}
	wxappSetting := model.WxappLoginSettings{}
	tx := db.Where("option_name = ?", "wxapp_login_setting").First(&option)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}
	value := option.OptionValue
	err := json.Unmarshal([]byte(value), &wxappSetting)

	if err != nil {
		resp.Error("解析时出错："+err.Error(), nil)
		return
	}

	resp.Success("获取成功！", wxappSetting)
	return
}
