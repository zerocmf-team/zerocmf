/**
** @创建时间: 2020/7/21 2:38 下午
** @作者　　: return
 */
package admin

import (
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strconv"
)

type AuthorizeController struct {
	rc controller.RestController
}

type tempAuthorize struct {
	Id         int     `json:"id"`
	RuleId     int     `json:"rule_id"`
	UniqueName string  `gorm:"type:varchar(30);comment:'唯一名称'" json:"unique_name"`
	ParentId   int     `gorm:"type:int(11);comment:'所属父类id';default:0" json:"parent_id"`
	Name       string  `gorm:"type:varchar(30);comment:'路由名称'" json:"name"`
	Path       string  `gorm:"type:varchar(100);comment:'路由路径'" json:"path"`
	Icon       string  `gorm:"type:varchar(30);comment:'图标名称'" json:"icon"`
	HideInMenu int     `gorm:"type:tinyint(3);comment:'菜单中隐藏';default:0" json:"hide_in_menu"`
	ListOrder  float64 `gorm:"type:float;comment:'排序';default:10000" json:"list_order"`
}

func (rest *AuthorizeController) Get(c *gin.Context) {
	var adminMenu []tempAuthorize
	result := cmf.Db.Find(&adminMenu)

	if result.RowsAffected == 0 {
		controller.RestController{}.Error(c, "暂无菜单,请联系管理员添加！", nil)
		return
	}

	results := rest.recursionMenu(adminMenu, 0,[]string{})
	rest.rc.Success(c, "获取成功！", results)
}

func (rest *AuthorizeController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	var adminMenu []tempAuthorize
	prefix := cmf.Conf().Database.Prefix
	result := cmf.Db.Debug().Table(prefix + "admin_menu m").
		Select("r.id as rule_id,m.id,m.unique_name,m.parent_id,m.name,m.path,m.icon,m.hide_in_menu,m.list_order").
		Joins("INNER JOIN  " + prefix + "auth_rule r ON m.unique_name = r.name").Scan(&adminMenu)

	if result.RowsAffected == 0 {
		controller.RestController{}.Error(c, "暂无菜单,请联系管理员添加！", nil)
		return
	}

	results := rest.recursionMenu(adminMenu, 0,[]string{})
	rest.rc.Success(c, "获取成功！", results)
}

func (rest *AuthorizeController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *AuthorizeController) Store(c *gin.Context) {
	rest.rc.Success(c, "操作成功Store", nil)
}

func (rest *AuthorizeController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}

type aResultStruct struct {
	Id       int           `json:"id"`
	RuleId   int           `json:"rule_id"`
	ParentId int           `json:"parent_id"`
	Title    string        `json:"title"`
	Key      int           `json:"key"`
	Children []interface{} `json:"children"`
}



func (rest *AuthorizeController) recursionMenu(menus []tempAuthorize, parentId int,tree []string) []aResultStruct {

	var results []aResultStruct

	tree = append(tree,strconv.Itoa(parentId))

	//key := strings.Join(tree,"-")

	for _, v := range menus {
		if parentId == v.ParentId {
			result := aResultStruct{
				Id:       v.Id,
				RuleId: v.RuleId,
				ParentId: v.ParentId,
				Title:    v.Name,
				Key:      v.RuleId,
			}

			routes := rest.recursionMenu(menus, v.Id,tree)
			itfRoutes := make([]interface{}, len(routes))
			for i, v := range routes {
				itfRoutes[i] = v
			}
			result.Children = itfRoutes
			results = append(results, result)
		}
	}
	return results
}
