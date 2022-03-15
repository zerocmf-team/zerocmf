/**
** @创建时间: 2021/12/21 13:21
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/casbin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"strconv"
	"strings"
)

type Role struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取所有角色列表
 * @Date 2021/12/21 13:42:33
 * @Param
 * @return
 **/

func (rest *Role) Get(c *gin.Context) {
	var query []string
	var queryArgs []interface{}
	//  用户状态
	status := c.Query("status")
	if status != "" {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, status)
	}
	// 名称模糊搜索
	name := c.Query("name")
	if name != "" {
		query = append(query, "name LIKE ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}
	var queryStr string
	queryStr = strings.Join(query, " AND ")
	current := c.DefaultQuery("current", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	intCurrent, _ := strconv.Atoi(current)
	intPageSize, _ := strconv.Atoi(pageSize)
	if intCurrent <= 0 {
		rest.Error(c, "当前页码需大于0！", nil)
		return
	}
	if intPageSize <= 0 {
		rest.Error(c, "每页数需大于0！", nil)
		return
	}
	db := util.GetDb(c)
	result, err := new(model.Role).Paginate(db, intCurrent, intPageSize, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", result)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查看单个角色信息
 * @Date 2021/12/23 12:12:52
 * @Param
 * @return
 **/

func (rest *Role) Show(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.Error(c, "角色id不能为空！", nil)
		return
	}
	db := util.GetDb(c)
	role := model.Role{}
	query := []string{"id = ?", "status = 1"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{id}
	err := role.Show(db, queryStr, queryArgs)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", role)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除一个或多个角色
 * @Date 2021/12/26 17:13:52
 * @Param
 * @return
 **/

func (rest *Role) Delete(c *gin.Context) {
	ids := c.QueryArray("ids")
	db := util.GetDb(c)
	role := model.Role{}
	e, err := casbin.NewEnforcer("")
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	if len(ids) == 0 {
		id := c.Param("id")
		if id == "" {
			rest.Error(c, "id不能为空！", nil)
			return
		}

		tx := db.Where("id = ?", id).First(&role)

		if util.IsDbErr(tx) != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		// 删除对应的角色关系
		e.DeleteRole(strconv.Itoa(role.Id))
		if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
			rest.Error(c, "删除失败！", err.Error())
			return
		}
	} else {
		if err := db.Where("id IN (?)", ids).Delete(&role).Error; err != nil {
			rest.Error(c, "删除失败！", nil)
			return
		}

		for _, v := range ids {
			e.DeleteRole(v)
		}

	}

	rest.Success(c, "删除成功！", nil)
}
