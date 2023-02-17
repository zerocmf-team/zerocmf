package adminMenu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"
	"zerocmf/service/user/rpc/types/user"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 行转树结构体
 * @Date 2021/11/30 12:52:26
 * @Param
 * @return
 **/

type routers struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Locale     string        `json:"locale"`
	Index      string        `json:"index"`
	Path       string        `json:"path"`
	Icon       string        `json:"icon"`
	HideInMenu int           `json:"hideInMenu"`
	ListOrder  float64       `json:"list_order"`
	CreateAt   int64         `json:"create_at"`
	CreateTime string        `json:"create_time"`
	Routes     []interface{} `json:"routes,omitempty"`
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取当前用户可访问的菜单
 * @Date 2022/3/13 19:18:49
 * @Param
 * @return
 **/

func (l *GetLogic) Get() (resp *types.Response) {
	resp = new(types.Response)
	c := l.svcCtx
	userRpc := c.UserRpc
	var menus []model.AdminMenu
	db := c.Db
	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)
	if tx.RowsAffected == 0 {
		resp.Error("暂无菜单，请和联系管理员添加！", nil)
		return
	}

	userId, _ := l.svcCtx.Get("userId")

	databaseReply, err := userRpc.Database(l.ctx, &user.DatabaseRequest{})
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	dbConf := database.Database{}
	copier.Copy(&dbConf, &databaseReply)

	//	获取当前用户的全部角色
	e, err := dbConf.NewEnforcer("")
	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var menusResult = make([]model.AdminMenu, 0)
	var access bool
	for _, v := range menus {

		path := v.Path
		if v.ParentId == 0 {
			path = v.Path
		}
		access, err = e.Enforce(userId, path, "*")
		if err != nil {
			resp.Error("系统出错", err.Error())
			// panic(err.Error())
		}
		if access {
			menusResult = append(menusResult, v)
		}
	}

	results := recursionMenu(menusResult, 0, "", "")
	if len(results) == 0 {
		results = make([]routers, 0)
	}

	resp.Success("获取成功！", results)
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归增加子菜单项
 * @Date 2021/11/30 12:50:24
 * @Param
 * @return
 **/

func recursionMenu(menus []model.AdminMenu, parentId int, parentIndex string, locale string) []routers {
	var routesResult []routers
	index := 0
	for _, v := range menus {

		if parentId == v.ParentId {
			iStr := strconv.Itoa(index)
			var curIndex string
			if parentIndex == "" {
				curIndex = iStr
			} else {
				curIndex = parentIndex + "-" + iStr
			}

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

			result := routers{
				Id:         v.Id,
				Index:      curIndex,
				Locale:     curLocale,
				Name:       v.Name,
				Path:       v.Path,
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
				CreateAt:   v.CreateAt,
				CreateTime: time.Unix(v.CreateAt, 0).Format(data.TimeLayout),
			}
			index++
			routes := recursionMenu(menus, v.Id, curIndex, curLocale)
			childRoutes := make([]interface{}, len(routes))
			for ri, rv := range routes {
				childRoutes[ri] = rv
			}
			result.Routes = childRoutes
			routesResult = append(routesResult, result)
		}
	}
	return routesResult
}
