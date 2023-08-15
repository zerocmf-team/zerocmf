package adminMenu

import (
	"context"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllMenusLogic {
	return &GetAllMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllMenusLogic) GetAllMenus() (resp *types.Response) {
	resp = new(types.Response)
	c := l.svcCtx
	var menus []model.AdminMenu
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)
	if tx.RowsAffected == 0 {
		resp.Error("暂无菜单，请和联系管理员添加！", nil)
		return
	}

	results := recursionMenu(menus, 0, "", "")
	if len(results) == 0 {
		results = make([]routers, 0)
	}

	resp.Success("获取成功！", results)
	return
}
