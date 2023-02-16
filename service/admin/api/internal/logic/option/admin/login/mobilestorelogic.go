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

type MobileStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMobileStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileStoreLogic {
	return &MobileStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MobileStoreLogic) MobileStore(req *types.MobileLoginReq) (resp types.Response) {
	c := l.svcCtx
	mobileSetting := model.MobileLoginSettings{}
	err := copier.Copy(&mobileSetting, &req)
	if err != nil {
		resp.Error("系统出错："+err.Error(), nil)
		return
	}

	mobileSettingValue, _ := json.Marshal(mobileSetting)

	db := c.Db
	tx := db.Model(&model.Option{}).Where("option_name = ?", "mobile_login_setting").Update("option_value", string(mobileSettingValue))
	if tx.Error != nil {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}

	resp.Success("修改成功！", req)
	return
}
