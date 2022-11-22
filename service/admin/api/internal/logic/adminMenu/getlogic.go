package adminMenu

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"
	"zerocmf/service/admin/model"
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
	Name       string        `gorm:"type:varchar(30);comment:'路由名称'" json:"name"`
	Index      string        `gorm:"type:varchar(255);comment:'分类层级关系路径'" json:"index"`
	Path       string        `gorm:"type:varchar(100);comment:'路由路径'" json:"path"`
	Icon       string        `gorm:"type:varchar(30);comment:'图标名称'" json:"icon"`
	HideInMenu int           `gorm:"type:tinyint(3);comment:'菜单中隐藏';default:0" json:"hideInMenu"`
	ListOrder  float64       `gorm:"type:float;comment:'排序';default:10000" json:"list_order"`
	CreateAt   int64         `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	CreateTime string        `gorm:"-" json:"create_time"`
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
	// todo: add your logic here and delete this line

	resp = new(types.Response)

	c := l.svcCtx
	var menus []model.AdminMenu
	db := c.Db
	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)
	if tx.RowsAffected == 0 {
		resp.Error("暂无菜单，请和联系管理员添加！", nil)
		return
	}

	userId, _ := l.svcCtx.Get("userId")
	database := database.Conf()

	//	获取当前用户的全部角色
	e, err := database.NewEnforcer("")
	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	//	存入casbin
	if err != nil {
		return
	}
	var menusResult = make([]model.AdminMenu, 0)
	var access bool
	for _, v := range menus {
		access, err = e.Enforce(userId, v.Path, "*")
		if err != nil {
			resp.Error("系统出错", err.Error())
			// panic(err.Error())
		}
		if access {
			menusResult = append(menusResult, v)
		}
	}
	rolePolicies := e.GetFilteredPolicy(0, userId.(string))

	if userId == "1" || len(rolePolicies) == 0 {
		menusResult = menus
	}
	results := recursionMenu(menusResult, 0, "")
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

func recursionMenu(menus []model.AdminMenu, parentId int, parentIndex string) []routers {
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

			result := routers{
				Id:         v.Id,
				Index:      curIndex,
				Name:       v.Name,
				Path:       v.Path,
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
				CreateAt:   v.CreateAt,
				CreateTime: time.Unix(v.CreateAt, 0).Format(data.TimeLayout),
			}
			index++
			routes := recursionMenu(menus, v.Id, curIndex)
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
