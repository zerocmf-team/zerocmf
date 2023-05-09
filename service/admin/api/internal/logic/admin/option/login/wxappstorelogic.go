package login

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxappStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxappStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxappStoreLogic {
	return &WxappStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxappStoreLogic) WxappStore(req *types.WxAppReq) (resp types.Response) {
	c := l.svcCtx
	wxappSetting := model.WxappLoginSettings{}
	err := copier.Copy(&wxappSetting, &req)
	if err != nil {
		resp.Error("系统出错："+err.Error(), nil)
		return
	}
	wxappSettingValue, _ := json.Marshal(wxappSetting)
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	tx := db.Model(&model.Option{}).Where("option_name = ?", "wxapp_login_setting").Update("option_value", string(wxappSettingValue))
	if tx.Error != nil {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}

	resp.Success("修改成功！", req)
	return
}
