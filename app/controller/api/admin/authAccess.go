/**
** @创建时间: 2020/8/17 9:00 上午
** @作者　　: return
** @描述　　:
 */
package admin

import (
	"fmt"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"strconv"
	"time"
)

type AuthAccessController struct {
	rc controller.RestController
}

// @Summary 根据角色id获取信息和允许权限
// @Description 根据角色id获取信息和允许权限
// @Tags 项目管理
// @Accept mpfd
// @Success 200 {object} model.ReturnData{data=[]model.ProjectPost}
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /admin/auth_access/:id [get]
// @Security Bearer Token
// @param Authorization header string true "Bearer Token"
func (rest *AuthAccessController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if rewrite.Id == 1 {
		rest.rc.Error(c,"超级管理员无法被编辑！",nil)
		return
	}

	role := model.Role{}
	err := cmf.Db.Where("id = ?",rewrite.Id).First(&role).Error

	if err != nil {
		rest.rc.Error(c,"查询角色失败！请联系管理员处理",nil)
		return
	}

	var access []model.AuthAccess

	cmf.Db.Where("role_id = ?",role.Id).Find(&access)

	var rule []int
	for _, v := range access{
		rule = append(rule,v.RuleId)
	}

	fmt.Println("rule",rule)

	result := make(map[string]interface{})

	result["name"] = role.Name
	result["remark"] = role.Remark
	result["access"] = rule

	rest.rc.Success(c,"获取成功！",result)

}

func (rest *AuthAccessController) Edit(c *gin.Context) {

	var rewrite struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if rewrite.Id == 1 {
		rest.rc.Error(c,"超级管理员无法被编辑！",nil)
		return
	}

	// 角色名称
	name := c.PostForm("name")
	if name == "" {
		rest.rc.Error(c,"名称不能为空！",nil)
		return
	}

	// 角色描述
	remark := c.PostForm("remark")
	if remark == "" {
		rest.rc.Error(c,"描述不能为空！",nil)
		return
	}

	// 角色授权列表
	roleAccess := c.PostFormArray("role_access")

	role := model.Role{
		Name: name,
		Remark: remark,
		CreateAt: time.Now().Unix(),
	}

	cmf.Db.Model(&role).Where("id = ?",rewrite.Id).Updates(role)

	// 查询当前存在的auth_access
	var access []model.AuthAccess
	cmf.Db.Where("role_id = ?",rewrite.Id).Find(&access)

	var arrTemp []interface{}
	for _,v := range roleAccess {
		arrTemp = append(arrTemp,v)
	}
	// 筛查出待删除的内容

	// 数据库中不包含的值
	for _,v := range access {
		ruleId := strconv.Itoa(v.RuleId)
		if !inArray(ruleId,arrTemp) {
			cmf.Db.Where("rule_id = ?",ruleId).Delete(&model.AuthAccess{})
		}
	}

	// 筛查出待添加的内容
	arrTemp = make([]interface{},0)
	for _,v := range access {
		ruleId := strconv.Itoa(v.RuleId)
		arrTemp = append(arrTemp,ruleId)
	}

	for _,v := range roleAccess {
		if !inArray(v,arrTemp) {
			ruleId,_ := strconv.Atoi(v)
			cmf.Db.Create(&model.AuthAccess{RoleId: rewrite.Id,RuleId: ruleId})
		}
	}

	rest.rc.Success(c,"操作成功！",nil)
}

func inArray(search interface{}, arr []interface{}) bool {

	for _, item := range arr {
		if search == item {
			return true
		}
	}
	return false
}

func (rest *AuthAccessController) Store(c *gin.Context) {

	// 角色名称
	name := c.PostForm("name")
	if name == "" {
		rest.rc.Error(c,"名称不能为空！",nil)
		return
	}

	// 角色描述
	remark := c.PostForm("remark")
	if remark == "" {
		rest.rc.Error(c,"描述不能为空！",nil)
		return
	}

	// 角色授权列表
	roleAccess := c.PostFormArray("role_access")

	role := model.Role{
		Name: name,
		Remark: remark,
		CreateAt: time.Now().Unix(),
	}

	cmf.Db.Where("name = ?",name).FirstOrCreate(&role)

	if role.Id == 0 {
		rest.rc.Error(c,"创建角色失败！请联系管理员",nil)
		return
	}

	for _,v := range roleAccess {
		ruleId,_ := strconv.Atoi(v)
		roleAccess := model.AuthAccess{
			RoleId: role.Id,
			RuleId: ruleId,
		}
		cmf.Db.Where("role_id = ? AND rule_id = ?",role.Id,ruleId).FirstOrCreate(&roleAccess)
	}

	rest.rc.Success(c, "操作成功！", role.Id)
}

