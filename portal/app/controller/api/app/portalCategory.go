/**
** @创建时间: 2021/12/13 19:16
** @作者　　: return
** @描述　　:
 */

package app

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/util"
	"strconv"
)

type PortalCategory struct {
	controller.Rest
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 根据当前id获取子类分类树
 * @Date 2022/2/10 14:8:15
 * @Param
 * @return
 **/

func (rest *PortalCategory) TreeList(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		rest.Error(c, "分类id不能为空！", nil)
		return
	}

	cidInt, err := strconv.Atoi(cid)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	db := util.GetDb(c)

	portalCategory := model.PortalCategory{
		ParentId: cidInt,
	}

	trees, err := portalCategory.ListWithTree(db)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", trees)

}

func (rest *PortalCategory) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	db := util.GetDb(c)
	data, err := new(model.PortalCategory).Show(db, "id = ? and delete_at = ?", []interface{}{rewrite.Id, 0})
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 根据id寻找父类id
 * @Date 2021/1/8 15:2:48
 * @Param
 * @return
 **/

func (rest *PortalCategory) GetTopId(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空", nil)
		return
	}
	idInt, err := strconv.Atoi(id)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	db := util.GetDb(c)
	topId, err := new(model.PortalCategory).GetTopId(db, idInt)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功！", gin.H{"top_id": topId})

}

func (rest *PortalCategory) Breadcrumb(c *gin.Context) {
	cid := c.Param("cid")
	cidInt, err := strconv.Atoi(cid)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if cidInt == 0 {
		rest.Error(c, "cid不能为空", nil)
		return
	}

	db := util.GetDb(c)
	breadcrumbs, err := new(model.PortalCategory).GetPrevious(db, cidInt)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	rest.Success(c, "获取成功", breadcrumbs)
}
