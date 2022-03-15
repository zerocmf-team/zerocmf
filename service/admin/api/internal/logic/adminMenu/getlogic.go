package adminMenu

import (
	"context"
	bsDb "gincmf/common/bootstrap/db"
	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"
	"gincmf/service/admin/model"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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
	Name       string        `gorm:"type:varchar(30);comment:'路由名称'" json:"name"`
	Path       string        `gorm:"type:varchar(100);comment:'路由路径'" json:"path"`
	Icon       string        `gorm:"type:varchar(30);comment:'图标名称'" json:"icon"`
	HideInMenu int           `gorm:"type:tinyint(3);comment:'菜单中隐藏';default:0" json:"hideInMenu"`
	ListOrder  float64       `gorm:"type:float;comment:'排序';default:10000" json:"list_order"`
	Routes     []interface{} `json:"routes"`
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

func (l *GetLogic) Get() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)

	c := l.svcCtx
	var menus []model.AdminMenu
	db := c.Db
	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)
	if tx.RowsAffected == 0 {
		result := c.Error("暂无菜单，请和联系管理员添加！", nil)
		copier.Copy(&resp, &result)
		return
	}

	userId, _ := l.svcCtx.Get("userId")
	database := bsDb.Database()

	//	获取当前用户的全部角色
	e, err := database.NewEnforcer("")
	//	存入casbin
	if err != nil {
		result := c.Error(err.Error(), nil)
		copier.Copy(&resp, &result)
		return
	}

	//	存入casbin
	if err != nil {
		return
	}
	var menusResult = make([]model.AdminMenu, 0)
	for _, v := range menus {
		access, err := e.Enforce(userId, v.Object,"*")
		if err != nil {
			panic(err)
		}
		if access {
			menusResult = append(menusResult, v)
		}
	}
	rolePolicies := e.GetFilteredPolicy(0, userId.(string))

	if userId == "1" || len(rolePolicies) == 0 {
		menusResult = menus
	}
	results := recursionMenu(menusResult, 0)
	if len(results) == 0 {
		results = make([]routers, 0)
	}

	result := c.Success("获取成功！", results)
	copier.Copy(&resp,&result)
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归增加子菜单项
 * @Date 2021/11/30 12:50:24
 * @Param
 * @return
 **/

func recursionMenu(menus []model.AdminMenu, parentId int) []routers {
	var routesResult []routers
	for _, v := range menus {
		if parentId == v.ParentId {
			result := routers{
				Name:       v.Name,
				Path:       v.Path,
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
			}

			routes := recursionMenu(menus, v.Id)
			childRoutes := make([]interface{}, len(routes))
			for i, v := range routes {
				childRoutes[i] = v
			}
			result.Routes = childRoutes
			routesResult = append(routesResult, result)
		}
	}
	return routesResult
}

