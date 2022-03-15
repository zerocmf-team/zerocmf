/**
** @创建时间: 2021/12/23 12:48
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
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AuthAccess struct {
	controller.Rest
}

func (rest *AuthAccess) Show(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "角色id不能为空！", nil)
	}
	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}
	db := util.GetDb(c)
	role := model.Role{}
	tx := db.Where("id = ? AND status = 1", id).First(&role)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			rest.Error(c, "该角色不存在或已删除！", nil)
			return
		}
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	id = strconv.Itoa(role.Id)

	// 获取全部角色策略
	roles := e.GetFilteredPolicy(0, id)
	result := make([]string, 0)
	for _, v := range roles {
		if len(v) > 1 {
			result = append(result, v[1])
		}
	}
	rest.Success(c, "获取成功！", result)
}

func (rest *AuthAccess) Store(c *gin.Context) {
	rest.Save(c, "")
}

func (rest *AuthAccess) Edit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.Error(c, "角色id不能为空！", nil)
	}
	rest.Save(c, id)
}

func (rest *AuthAccess) Save(c *gin.Context, editId string) {

	var form struct {
		Name       string   `json:"name"`
		Remark     string   `json:"remark"`
		RoleAccess []string `json:"role_access"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	db := util.GetDb(c)

	// 角色信息
	role := model.Role{
		Name:     form.Name,
		Remark:   form.Remark,
		CreateAt: time.Now().Unix(),
		Status:   1,
	}

	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if editId == "" {
		tx := db.Create(&role)
		if tx.Error != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		id := strconv.Itoa(role.Id)

		for _, v := range form.RoleAccess {
			e.AddPolicy(id, v,"*")
		}

	} else {
		roleItem := model.Role{}
		tx := db.Where("id = ? AND status = 1", editId).First(&roleItem)
		if tx.Error != nil {
			if tx.Error == gorm.ErrRecordNotFound {
				rest.Error(c, "该角色不存在或已删除！", nil)
				return
			}
			rest.Error(c, tx.Error.Error(), nil)
			return
		}
		role.Id = roleItem.Id
		tx = db.Save(&role)
		if tx.Error != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		id := strconv.Itoa(role.Id)

		// 新增修改策略

		// 获取全部角色策略
		roles := e.GetFilteredPolicy(0, id)
		existAccess := make([]string, 0)
		for _, v := range roles {
			if len(v) > 1 {
				existAccess = append(existAccess, v[1])
			}
		}

		/*
		*  新增：[1,2,3]
		*  原有：[3,4,5]
		*  筛选去除的规则：[4，5]
		*  筛选新增的规则：[1，2]

		* 新增：[1,2,3]
		* 原有：[]
		* 筛选去除的规则：[]
		* 筛选新增的规则：[1，2，3]

		* 新增：[]
		* 原有：[1，2，3]
		* 筛选去除的规则：[1，2，3]
		* 筛选新增的规则：[]
		 */

		alreadyDel := make([]string, 0)
		// 判断是否需要被删除
		for _, v := range existAccess {
			if util.ToLowerInArray(v, form.RoleAccess) == false {
				alreadyDel = append(alreadyDel, v)
			}
		}
		// 如果新增为空，则全部删除
		if len(form.RoleAccess) == 0 {
			alreadyDel = existAccess
		}

		alreadyAdd := make([]string, 0)
		for _, v := range form.RoleAccess {
			if util.ToLowerInArray(v, existAccess) == false {
				alreadyAdd = append(alreadyAdd, v)
			}
		}

		// 如果数据库不存在，则为新增
		if len(existAccess) == 0 {
			alreadyAdd = form.RoleAccess
		}

		// 开始删除策略
		rules := make([][]string, 0)
		for _, v := range alreadyDel {
			rules = append(rules, []string{id, v,"*"})
		}
		if len(rules) > 0 {
			e.RemovePolicies(rules)
		}

		// 开始新增策略
		rules = make([][]string, 0)
		for _, v := range alreadyAdd {
			rules = append(rules, []string{id, v,"*"})
		}
		if len(rules) > 0 {
			e.AddPolicies(rules)
		}

	}

	rest.Success(c, "操作成功！", role)

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除角色和权限记录
 * @Date 2021/12/24 13:12:10
 * @Param
 * @return
 **/

func (rest *AuthAccess) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.Error(c, "角色id不能为空！", nil)
	}

	db := util.GetDb(c)
	role := model.Role{}
	tx := db.Where("id = ? AND status = 1", id).First(&role)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			rest.Error(c, "该角色不存在或已删除！", nil)
			return
		}
		rest.Error(c, tx.Error.Error(), nil)
		return
	}

	//	删除casbin权限


}
