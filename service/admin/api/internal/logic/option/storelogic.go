package option

import (
	"context"
	"encoding/json"
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

func (l *StoreLogic) Store(req types.OptionReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	c := l.svcCtx
	siteInfoValue, _ := json.Marshal(req)

	db := c.Db
	tx := db.Model(&model.Option{}).Where("option_name = ?", "site_info").Update("option_value", string(siteInfoValue))
	if tx.Error != nil {
		resp.Error("系统出错："+tx.Error.Error(), nil)
		return
	}

	resp.Success("修改成功", req)
	return
}
