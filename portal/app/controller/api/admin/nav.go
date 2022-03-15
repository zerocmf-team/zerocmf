/**
** @创建时间: 2021/1/3 8:29 下午
** @作者　　: return
** @描述　　:
 */

package admin

import (
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/bootstrap/controller"
	"github.com/gincmf/bootstrap/paginate"
	"github.com/gincmf/bootstrap/util"
)

type Nav struct {
	controller.Rest
}

func (rest *Nav) Get(c *gin.Context) {
	db := util.GetDb(c)
	current, pageSize, err := new(paginate.Paginate).Default(c)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	data, err := new(model.Nav).Get(db, current, pageSize, "", nil)
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	rest.Success(c, "获取成功！", data)
}

func (rest *Nav) Edit(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.Save(c, rewrite.Id)

}

func (rest *Nav) Store(c *gin.Context) {
	rest.Save(c, 0)
}

func (rest *Nav) Save(c *gin.Context, editId int) {

	var form struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		Remark string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, err.Error())
		return
	}

	db := util.GetDb(c)

	nav := model.Nav{
		Key:    form.Key,
		Name:   form.Name,
		Remark: form.Remark,
	}

	// 新增
	if editId == 0 {
		db.Create(&nav)
	} else {
		queryNav := model.Nav{}
		tx := db.Where("id = ?", editId).First(&queryNav)
		if tx.Error != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}
		nav.Id = queryNav.Id
		db.Save(&nav)
	}
	rest.Success(c, "操作成功！", nil)
}

func (rest *Nav) Delete(c *gin.Context) {
	rest.Success(c, "操作成功Delete", nil)
}
