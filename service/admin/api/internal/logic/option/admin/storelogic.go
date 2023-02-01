package admin

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreLogic {
	return StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.OptionReq) (resp types.Response) {
	c := l.svcCtx
	siteInfo := model.SiteInfo{}
	err := copier.Copy(&siteInfo, &req)
	if err != nil {
		resp.Error("系统出错："+err.Error(), nil)
		return
	}

	siteInfoValue, _ := json.Marshal(siteInfo)

	db := c.Db
	tx := db.Model(&model.Option{}).Where("option_name = ?", "site_info").Update("option_value", string(siteInfoValue))
	if tx.Error != nil {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}

	resp.Success("修改成功", req)
	return
}
