package navItem

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
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

func (l *StoreLogic) Store(req *types.NavItemSaveReq) (resp types.Response) {
	resp = save(l.svcCtx, req)
	return
}

func save(c *svc.ServiceContext, req *types.NavItemSaveReq) (resp types.Response) {

	if req.NavId == 0 {
		resp.Error("导航id不能为空！", nil)
		return
	}

	if req.Name == "" {
		resp.Error("导航项名称不能为空！", nil)
		return
	}

	db := c.Db

	editId := req.Id

	navItem := model.NavItem{}
	copier.Copy(&navItem, &req)

	msg := ""
	if editId > 0 {
		tempNavItem := model.NavItem{}
		tx := db.Where("id = ?", editId).First(&tempNavItem)

		if tx.Error != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}

		navItem.Id = tempNavItem.Id
		db.Save(&navItem)
		msg = "保存成功！"
	} else {
		db.Create(&navItem)
		msg = "创建成功！"
	}

	resp.Success(msg, navItem)

	return
}
