/**
Desc: 添加菜单
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:38:05
*/

package adminMenu

import (
	"context"
	"zerocmf/service/lowcode/model"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

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

/**
Desc: 菜单类型：目录，页面，插件，按钮
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:39:08
*/

func (l *StoreLogic) Store(req *types.AdminMenuSaveReq) (resp types.Response) {
	return save(l.svcCtx, req)
}

/**
Desc: 新增和修改
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:39:08
*/

func save(c *svc.ServiceContext, req *types.AdminMenuSaveReq) (resp types.Response) {
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	adminMenu := model.AdminMenu{}
	// 0 => 表单/页面 1 => 组件 2=> 按钮
	var formId primitive.ObjectID
	mType := req.MenuType
	if mType == 0 {
		if req.FormId == nil {
			resp.Error("表单不能为空！", nil)
			return
		}

		formId, err = primitive.ObjectIDFromHex(*req.FormId)
		if err != nil {
			resp.Error("非法表单id", err.Error())
			return
		}

		// todo 校验表单的合法性，表单是否真实存在

		adminMenu.FormId = &formId
	}

	err = copier.Copy(&adminMenu, req)
	if err != nil {
		resp.Error(" copier.Copy错误！", err.Error())
		return
	}

	parentId := req.ParentId
	if parentId != nil && *parentId != "" {
		adminMenu.ParentId, err = primitive.ObjectIDFromHex(*parentId)
		if err != nil {
			resp.Error("非法parentId", err.Error())
			return
		}
	}

	err = adminMenu.Save(db, nil)
	if err != nil {
		resp.Error("新增失败", err.Error())
		return
	}
	resp.Success("新增成功", adminMenu)
	return
}
