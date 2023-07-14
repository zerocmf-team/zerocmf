package adminMenu

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.MenuReq) (resp *types.Response) {
	resp = MenuSave(l.svcCtx, req)
	return
}

func MenuSave(svcCtx *svc.ServiceContext, req *types.MenuReq) (resp *types.Response) {
	resp = new(types.Response)
	id := req.Id
	db := svcCtx.Db

	menu := model.AdminMenu{}
	copier.Copy(&menu, &req)

	var tx *gorm.DB
	if id > 0 {
		tx = db.Save(&menu)
	} else {
		menu.CreateAt = time.Now().Unix()
		tx = db.Create(&menu)
	}
	if tx.Error != nil {
		resp.Error("操作失败", tx.Error)
		return
	}

	resp.Success("操作成功！", nil)
	return
}
