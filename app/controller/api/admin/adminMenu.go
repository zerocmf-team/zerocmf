/**
** @创建时间: 2020/7/18 7:07 下午
** @作者　　: return
 */
package admin

import (
	"fmt"
	"gincmf/app/model"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
)

type MenuController struct{}

func (rest *MenuController) Get(c *gin.Context) {
	var adminMenu []model.AdminMenu
	result := cmf.Db.Where("path <> ?", "").Order("list_order").Find(&adminMenu)

	if result.RowsAffected == 0 {
		controller.RestController{}.Error(c, "暂无菜单,请联系管理员添加！", nil)
		return
	}

	authAccessRule := util.AuthAccess(c)
	fmt.Println("authAccess",authAccessRule)

	var showMenu []model.AdminMenu
	for _,v := range adminMenu{
		if rest.inMap(v.UniqueName,authAccessRule){
			showMenu = append(showMenu,v)
		}
	}

	if len(authAccessRule) == 0 {
		showMenu = adminMenu
	}

	fmt.Println("showMenu",showMenu)

	results := rest.recursionMenu(c,showMenu, 0)
	controller.RestController{}.Success(c, "获取成功！", results)

}

type resultStruct struct {
	Name       string        `gorm:"type:varchar(30);comment:'路由名称'" json:"name"`
	Path       string        `gorm:"type:varchar(100);comment:'路由路径'" json:"path"`
	Icon       string        `gorm:"type:varchar(30);comment:'图标名称'" json:"icon"`
	HideInMenu int           `gorm:"type:tinyint(3);comment:'菜单中隐藏';default:0" json:"hideInMenu"`
	ListOrder  float64       `gorm:"type:float;comment:'排序';default:10000" json:"list_order"`
	Routes     []interface{} `json:"routes"`
}

func (rest *MenuController) inMap(s string,target []model.AuthAccessRule) (result bool){
	fmt.Println(s,target)
	for _,v := range target{
		if s == v.Name {
			return true
		}
	}
	return  false
}

func (rest *MenuController) recursionMenu(c *gin.Context,menus []model.AdminMenu, parentId int) []resultStruct {

	var results []resultStruct
	for _, v := range menus {
		if parentId == v.ParentId {
			result := resultStruct{
				Name:       v.Name,
				Path:       v.Path,
				Icon:       v.Icon,
				HideInMenu: v.HideInMenu,
				ListOrder:  v.ListOrder,
			}

			routes := rest.recursionMenu(c,menus, v.Id)
			childRoutes := make([]interface{}, len(routes))
			for i, v := range routes {
				childRoutes[i] = v
			}
			result.Routes = childRoutes
			results = append(results, result)
		}
	}
	return results
}
