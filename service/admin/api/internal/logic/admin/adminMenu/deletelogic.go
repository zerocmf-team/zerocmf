package adminMenu

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IdReq) (resp *types.Response) {
	resp = new(types.Response)
	db := l.svcCtx.Db
	id := req.Id

	adminMenu := model.AdminMenu{}
	tx := db.Where("id = ?", id).First(&adminMenu)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			resp.Error("该菜单不存在或已被删除！", nil)
			return
		}
		resp.Error("系统错误", nil)
		return
	}

	var count int64 = 0
	tx = db.Where("parent_id", id).Find(&adminMenu).Count(&count)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	if count > 0 {
		resp.Error("请先删除子集菜单！", nil)
		return
	}

	tx = db.Where("id", id).Delete(&adminMenu)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
