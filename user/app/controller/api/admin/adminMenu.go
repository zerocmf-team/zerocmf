/**
** @创建时间: 2021/11/29 12:51
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/model"
	"github.com/gincmf/bootstrap/util"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 显示站点菜单
 * @Date 2021/11/29 12:51:58
 * @Param
 * @return
 **/

type Menu struct {
	controller.Rest
}

func (rest *Menu) Get(c *gin.Context) {

	db := util.GetDb(c)

	var menus []model.AdminMenu

	tx := db.Where("path <> ?", "").Order("list_order, id").Find(&menus)

	if tx.RowsAffected == 0 {
		rest.Error(c, "暂无管理员，亲和联系管理员添加！", nil)
		return
	}

	results := rest.recursionMenu(menus, 0)
	rest.Success(c, "获取成功！", results)

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

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归增加子菜单项
 * @Date 2021/11/30 12:50:24
 * @Param
 * @return
 **/

func (rest *Menu) recursionMenu(menus []model.AdminMenu, parentId int) []routers {
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

			routes := rest.recursionMenu(menus, v.Id)
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
