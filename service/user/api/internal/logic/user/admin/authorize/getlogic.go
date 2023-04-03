package authorize

import (
	"context"
	"github.com/jinzhu/copier"
	"strings"
	"zerocmf/service/admin/model"
	"zerocmf/service/admin/rpc/types/admin"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

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
	Locale   string       `json:"locale"`
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

func (l *GetLogic) Get() (resp types.Response) {

	c := l.svcCtx
	adminRpc := c.AdminRpc
	getMenus, err := adminRpc.GetMenus(l.ctx, &admin.AdminMenuReq{})
	if err != nil {
		resp.Error(err.Error(), nil)
	}

	menuData := getMenus.GetData()
	var menus []model.AdminMenu

	copier.Copy(&menus, &menuData)

	results := recursionMenu(menus, 0, "")
	resp.Success("获取成功！", results)
	return
}

func recursionMenu(menus []model.AdminMenu, parentId int, locale string) []authorizes {
	var result []authorizes
	for _, v := range menus {

		var curLocale string
		if locale == "" {
			curLocale = "menu." + v.Name
		} else {
			if strings.HasPrefix(locale, "/") {
				curLocale = v.Name
			} else {
				curLocale = locale + "." + v.Name
			}
		}

		if parentId == v.ParentId {
			item := authorizes{
				Id:       v.Id,
				ParentId: v.ParentId,
				Title:    v.Name,
				Locale:   curLocale,
				Key:      v.Path,
			}
			routes := recursionMenu(menus, v.Id, curLocale)
			item.Children = routes
			result = append(result, item)
		}
	}
	return result
}
