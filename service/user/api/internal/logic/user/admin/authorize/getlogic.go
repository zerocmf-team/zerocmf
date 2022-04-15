package authorize

import (
	"context"
	"gincmf/service/admin/model"
	"gincmf/service/admin/rpc/types/admin"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type authorizes struct {
	Id       int          `json:"id"`
	ParentId int          `json:"parent_id"`
	Title    string       `json:"title"`
	Key      string       `json:"key"`
	Children []authorizes `json:"children"`
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp *types.Response, err error) {
	resp = new(types.Response)
	c := l.svcCtx

	adminRpc := c.AdminRpc
	getMenus, err := adminRpc.GetMenus(l.ctx,&admin.AdminMenuReq{})
	if err != nil {
		resp.Error(err.Error(), nil)
	}

	menuData := getMenus.GetData()
	var menus []model.AdminMenu

	copier.Copy(&menus,&menuData)

	results := recursionMenu(menus, 0)
	resp.Success("获取成功！", results)
	return
}

func recursionMenu(menus []model.AdminMenu, parentId int) []authorizes {
	var result []authorizes
	for _, v := range menus {
		if parentId == v.ParentId {
			item := authorizes{
				Id:       v.Id,
				ParentId: v.ParentId,
				Title:    v.Name,
				Key:      v.Object,
			}
			routes := recursionMenu(menus, v.Id)
			item.Children = routes
			result = append(result, item)
		}
	}
	return result
}
