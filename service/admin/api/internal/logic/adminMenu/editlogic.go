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

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.MenuReq) (resp *types.Response) {
	c := l.svcCtx
	db := c.Db
	menu := model.AdminMenu{}
	id := req.Id
	if id <= 0 {
		resp.Error("参数不合法", nil)
		return
	}

	tx := db.Where("id = ?", id).First(&menu)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			resp.Error("该菜单不存在或已被删除！", nil)
			return
		}
		resp.Error("系统错误", nil)
		return
	}
	resp = MenuSave(l.svcCtx, req)
	return
}
