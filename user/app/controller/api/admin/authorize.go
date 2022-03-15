/**
** @创建时间: 2021/12/22 19:32
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

type Authorize struct {
	controller.Rest
}

func (rest *Authorize) Get(c *gin.Context) {
	// 查询菜单，并按树形结构显示
	// 验证当前
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

type authorizes struct {
	Id       int          `json:"id"`
	ParentId int          `json:"parent_id"`
	Title    string       `json:"title"`
	Key      string       `json:"key"`
	Children []authorizes `json:"children"`
}

func (rest *Authorize) recursionMenu(menus []model.AdminMenu, parentId int) []authorizes {
	var result []authorizes
	for _, v := range menus {
		if parentId == v.ParentId {
			item := authorizes{
				Id:       v.Id,
				ParentId: v.ParentId,
				Title:    v.Name,
				Key:      v.Object,
			}
			routes := rest.recursionMenu(menus, v.Id)
			item.Children = routes
			result = append(result, item)
		}
	}
	return result
}

