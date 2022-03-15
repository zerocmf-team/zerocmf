/**
** @创建时间: 2020/7/19 7:24 下午
** @作者　　: return
 */
package admin

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strconv"
	"strings"
)

type RoleController struct {
	rc controller.RestController
}

func (rest *RoleController) Get(c *gin.Context) {

	var role []model.Role

	var query []string
	var queryArgs []interface{}

	//  用户状态
	status  := c.Query("status")
	if status != "" {
		query = append( query,"status = ?")
		queryArgs = append(queryArgs,status)
	}

	// 名称模糊搜索
	name := c.Query("name")
	if name != "" {
		query = append( query,"name LIKE ?")
		queryArgs = append(queryArgs,"%"+name+"%")
	}

	var queryStr interface{}
	queryStr = strings.Join(query," AND ")

	current := c.DefaultQuery("current", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	intCurrent, _ := strconv.Atoi(current)
	intPageSize, _ := strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		rest.rc.Error(c, "当前页码需大于0！", nil)
		return
	}

	if intPageSize <= 0 {
		rest.rc.Error(c, "每页数需大于0！", nil)
		return
	}

	var total int64 = 0

	cmf.Db.Where(queryStr, queryArgs...).Find(&role).Count(&total)
	result := cmf.Db.Limit(intPageSize).Where(queryStr, queryArgs...).Offset((intCurrent - 1) * intPageSize).Find(&role)

	if result.RowsAffected == 0 {
		rest.rc.Error(c, "该页码内容不存在！", role)
		return
	}

	paginationData := &model.Paginate{Data: role, Current: current, PageSize: pageSize, Total: total}
	if len(role) == 0 {
		paginationData.Data = make([]string, 0)
	}

	rest.rc.Success(c, "获取成功", paginationData)
}

func (rest *RoleController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *RoleController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *RoleController) Store(c *gin.Context) {
	rest.rc.Success(c, "操作成功Store", nil)
}

func (rest *RoleController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
