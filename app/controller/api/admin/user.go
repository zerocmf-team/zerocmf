/**
** @创建时间: 2020/8/18 8:48 上午
** @作者　　: return
** @描述　　:
 */
package admin

import (
	"fmt"
	"gincmf/app/model"
	gUtil "gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/util"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	rc controller.RestController
}

func (rest *UserController) Get(c *gin.Context) {

	var user []model.User

	query := []string{"user_type = ?"}
	queryArgs := []interface{}{"1"}

	userLogin := c.Query("user_login")
	if userLogin != "" {
		query = append(query, "user_login LIKE ?")
		queryArgs = append(queryArgs, "%"+userLogin+"%")
	}

	userNickname := c.Query("user_nickname")
	if userNickname != "" {
		query = append(query, "user_nickname like ?")
		queryArgs = append(queryArgs, "%"+userNickname+"%")
	}

	userEmail := c.Query("user_email")
	if userEmail != "" {
		query = append(query, "user_email like ?")
		queryArgs = append(queryArgs, "%"+userEmail+"%")
	}

	queryStr := strings.Join(query, " AND ")

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

	cmf.Db.Where(queryStr, queryArgs...).Find(&user).Count(&total)
	result := cmf.Db.Where(queryStr, queryArgs...).Limit(intPageSize).Offset((intCurrent - 1) * intPageSize).Find(&user)

	if result.RowsAffected == 0 {
		rest.rc.Error(c, "该页码内容不存在！", nil)
		return
	}

	type temResult struct {
		model.User
		LastLoginTime string `json:"last_login_time"`
		CreateTime    string `json:"create_time"`
	}
	var tempResult []temResult
	for _, v := range user {
		var (
			lastLoginTime string
			createTime    string
		)
		if v.LastLoginAt == 0 {
			lastLoginTime = "0"
		} else {
			lastLoginTime = time.Unix(v.LastLoginAt, 0).Format("2006-01-02 15:04:05")
		}
		if v.CreateAt == 0 {
			createTime = "0"
		} else {
			createTime = time.Unix(v.CreateAt, 0).Format("2006-01-02 15:04:05")
		}
		tempResult = append(tempResult, temResult{User: v, LastLoginTime: lastLoginTime, CreateTime: createTime})
	}

	paginationData := &model.Paginate{Data: tempResult, Current: current, PageSize: pageSize, Total: total}
	if len(tempResult) == 0 {
		paginationData.Data = make([]string, 0)
	}

	rest.rc.Success(c, "获取成功！", paginationData)
}

func (rest *UserController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	query := "id = ? AND user_type = ?"
	queryArgs := []interface{}{rewrite.Id, "1"}

	user := model.User{}
	cmf.Db.Where(query, queryArgs...).First(&user)

	type resultStruct struct {
		model.User
		RoleIds   []int  `json:"role_ids"`
	}

	var roleUser []model.RoleUser
	cmf.Db.Where("user_id = ?", user.Id).Find(&roleUser)

	var role []int
	for _, v := range roleUser {
		role = append(role, v.RoleId)
	}

	result := resultStruct{
		User:user,
		RoleIds:   role,
	}

	rest.rc.Success(c, "获取成功！", result)
}

func (rest *UserController) Edit(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	userLogin := c.PostForm("user_login")
	if userLogin == "" {
		rest.rc.Error(c, "用户名不能为空！", nil)
		return
	}

	password := c.PostForm("user_pass")

	email := c.PostForm("user_email")

	mobile := c.PostForm("mobile")

	realName := c.PostForm("user_realname")

	roleIds := c.PostFormArray("role_ids")
	if len(roleIds) <= 0 {
		rest.rc.Error(c, "角色至少选择一项！", nil)
		return
	}

	departmentId := c.PostForm("department_id")
	if departmentId == "" {
		rest.rc.Error(c, "所在部门不能为空！", nil)
		return
	}
	departmentIdInt, _ := strconv.Atoi(departmentId)

	user := model.User{}

	result := cmf.Db.Where("user_login = ?", userLogin).First(&user)
	if result.RowsAffected == 0 {
		rest.rc.Error(c, "用户不存在！", nil)
		return
	}

	user.UserType = 1
	user.Mobile = mobile
	user.UserRealName = realName
	user.UserLogin = userLogin
	user.UserEmail = email
	user.DepartmentId = departmentIdInt
	user.UpdateAt = time.Now().Unix()
	user.UserStatus = 1

	if user.UserPass != "" {
		user.UserPass = util.GetMd5(password)
	}

	err := cmf.Db.Save(&user).Error
	if err != nil {
		rest.rc.Error(c, "更新用户出错，请联系管理员！！", nil)
		return
	}

	// 删除原来角色
	cmf.Db.Where("user_id = ?", rewrite.Id).Delete(&model.RoleUser{})

	// 存入用户角色
	for _, v := range roleIds {
		roleId, _ := strconv.Atoi(v)
		roleUser := model.RoleUser{
			RoleId: roleId,
			UserId: rewrite.Id,
		}
		cmf.Db.Create(&roleUser)
	}

	rest.rc.Success(c, "更新成功！", nil)
}

func (rest *UserController) Store(c *gin.Context) {

	userLogin := c.PostForm("user_login")
	if userLogin == "" {
		rest.rc.Error(c, "用户名不能为空！", nil)
		return
	}

	password := c.PostForm("user_pass")
	if password == "" {
		rest.rc.Error(c, "密码不能为空！", nil)
		return
	}

	email := c.PostForm("user_email")

	mobile := c.PostForm("mobile")
	realName := c.PostForm("user_realname")


	roleIds := c.PostFormArray("role_ids")
	if len(roleIds) <= 0 {
		rest.rc.Error(c, "角色至少选择一项！", nil)
		return
	}

	departmentId := c.PostForm("department_id")
	if departmentId == "" {
		rest.rc.Error(c, "所在部门不能为空！", nil)
		return
	}
	departmentIdInt, _ := strconv.Atoi(departmentId)

	user := model.User{
		UserType:     1,
		CreateAt:     time.Now().Unix(),
		Mobile: mobile,
		UserRealName: realName,
		UserLogin:    userLogin,
		UserPass:     util.GetMd5(password),
		UserEmail:    email,
		DepartmentId: departmentIdInt,
		UserStatus:   1,
	}

	result := cmf.Db.Where("user_login = ?", userLogin).First(&model.User{})

	if result.RowsAffected > 0 {
		rest.rc.Error(c, "用户已存在！", nil)
		return
	}

	err := cmf.Db.Create(&user).Error
	if err != nil {
		rest.rc.Error(c, "创建用户出错，请联系管理员！！", nil)
		return
	}

	// 存入用户角色

	userId := user.Id
	fmt.Println("userId",userId)

	for _, v := range roleIds {
		roleId, _ := strconv.Atoi(v)
		roleUser := model.RoleUser{
			RoleId: roleId,
			UserId: userId,
		}
		cmf.Db.Create(&roleUser)
	}

	rest.rc.Success(c, "操作成功！", user)
}

func (rest *UserController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}

func (rest *UserController) CurrentUser(c *gin.Context) {
	// 获取当前用户
	var currentUser = gUtil.CurrentUser(c)

	type temp struct {
		model.User
	}

	result:= temp{
		User:*currentUser,
	}

	controller.RestController{}.Success(c, "获取成功", result)
}
