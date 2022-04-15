package logic

import (
	"context"
	"errors"
	"gincmf/service/admin/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"gincmf/service/admin/rpc/internal/svc"
	"gincmf/service/admin/rpc/types/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenusLogic {
	return &GetMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenusLogic) GetMenus(in *admin.AdminMenuReq) (*admin.AdminMenuReply, error) {

	c := l.svcCtx
	db := c.Db

	// 获取全部菜单信息
	var menus []model.AdminMenu
	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return &admin.AdminMenuReply{}, errors.New("请联系管理员先添加菜单！")
		}
		return &admin.AdminMenuReply{}, tx.Error
	}

	var data []*admin.AdminMenu

	copier.Copy(&data, &menus)

	return &admin.AdminMenuReply{
		Data: data,
	}, nil
}
