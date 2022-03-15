/**
** @创建时间: 2021/12/20 12:45
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
	"time"
)

type Account struct {
	controller.Rest
}

func (rest *Account) Get(c *gin.Context) {

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
		rest.Error(c, "当前页码需大于0！", nil)
		return
	}

	if intPageSize <= 0 {
		rest.Error(c, "每页数需大于0！", nil)
		return
	}

	db := util.GetDb(c)
	result, err := new(model.User).Paginate(db, intCurrent, intPageSize, queryStr, queryArgs)

	if err != nil {
		rest.Error(c, "获取失败："+err.Error(), nil)
	}

	rest.Success(c, "获取成功！", result)
}

func (rest *Account) Show(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		rest.Error(c, "id不能为空！", nil)
		return
	}
	db := util.GetDb(c)
	user := model.User{}
	tx := db.Where("id = ? AND user_status = 1", []interface{}{id}).First(&user)
	if util.IsDbErr(tx) != nil {
		rest.Error(c, tx.Error.Error(), nil)
		return
	}
	if user.Id == 0 {
		rest.Error(c, "该用户不存在！", nil)
		return
	}

	userId := strconv.Itoa(user.Id)

	//	获取当前用户的全部角色
	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	roles, err := e.GetRolesForUser(userId)

	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	var result struct {
		model.User
		Roles []string `json:"roles"`
	}

	result.User = user
	result.Roles = roles

	rest.Success(c, "获取成功！", result)

}

func (rest *Account) Store(c *gin.Context) {
	rest.Save(c, "0")
}

func (rest *Account) Edit(c *gin.Context) {
	id := c.Param("id")
	rest.Save(c, id)
}

func (rest *Account) Save(c *gin.Context, editId string) {

	userId, _ := strconv.Atoi(editId)

	var form struct {
		UserLogin    string   `json:"user_login"`
		UserPass     string   `json:"user_pass"`
		UserEmail    string   `json:"user_email"`
		Mobile       string   `json:"mobile"`
		UserRealname string   `json:"user_realname"`
		RoleIds      []string `json:"role_ids"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if len(form.RoleIds) <= 0 {
		rest.Error(c, "至少选择一项角色！", nil)
		return
	}

	user := model.User{
		UserType:     1,
		CreateAt:     time.Now().Unix(),
		Mobile:       form.Mobile,
		UserRealName: form.UserRealname,
		UserLogin:    form.UserLogin,
		UserPass:     util.GetMd5(form.UserPass),
		UserEmail:    form.UserEmail,
		UserStatus:   1,
	}

	db := util.GetDb(c)

	// 存入用户角色
	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		rest.Error(c, err.Error(), nil)
		return
	}

	if editId == "" {

		currentUser := model.User{}
		tx := db.Where("user_login = ?", form.UserLogin).First(&currentUser)

		if util.IsDbErr(tx) != nil {
			rest.Error(c, tx.Error.Error(), nil)
			return
		}

		if currentUser.Id > 0 {
			rest.Error(c, "该用户已存在！", nil)
			return
		}

		tx = db.Create(&user)
		if tx.Error != nil {
			rest.Error(c, "创建用户出错，请联系管理员！", tx.Error)
			return
		}
		userId := strconv.Itoa(currentUser.Id)
		roleIds := form.RoleIds
		rules := make([][]string, 0)
		for _, v := range roleIds {
			rules = append(rules, []string{userId, v})
		}
		if len(rules) > 0 {
			e.AddGroupingPolicies(rules)
		}
	} else {

		editUser := model.User{}
		tx := db.Where("id = ?", editId).First(&editUser)

		if editUser.Id > 0 && editUser.UserLogin != form.UserLogin {
			currentUser := model.User{}
			tx := db.Where("user_login = ?", form.UserLogin).First(&currentUser)
			if util.IsDbErr(tx) != nil {
				rest.Error(c, tx.Error.Error(), nil)
				return
			}
			if currentUser.Id > 0 {
				rest.Error(c, "该登录名已存在！", nil)
				return
			}
		}

		if form.UserPass == "" {
			user.UserPass = util.GetMd5(form.UserPass)
		}
		user.Id = userId
		tx = db.Save(&user)

		if tx.Error != nil {
			rest.Error(c, "创建用户出错，请联系管理员！", tx.Error)
			return
		}

		roles, err := e.GetRolesForUser(editId)
		if err != nil {
			rest.Error(c, err.Error(), nil)
			return
		}

		alreadyDel := make([]string, 0)
		// 判断是否需要被删除
		for _, v := range roles {
			if util.ToLowerInArray(v, form.RoleIds) == false {
				alreadyDel = append(alreadyDel, v)
			}
		}
		// 如果新增为空，则全部删除
		if len(form.RoleIds) == 0 {
			alreadyDel = roles
		}

		alreadyAdd := make([]string, 0)
		for _, v := range form.RoleIds {
			if util.ToLowerInArray(v, roles) == false {
				alreadyAdd = append(alreadyAdd, v)
			}
		}

		// 如果数据库不存在，则为新增
		if len(roles) == 0 {
			alreadyAdd = form.RoleIds
		}

		// 开始删除策略
		rules := make([][]string, 0)
		for _, v := range alreadyDel {
			rules = append(rules, []string{editId,v})
		}
		if len(rules) > 0 {
			e.RemoveGroupingPolicies(rules)
		}

		// 开始新增策略
		rules = make([][]string, 0)
		for _, v := range alreadyAdd {
			rules = append(rules, []string{editId,v})
		}
		if len(rules) > 0 {
			e.AddGroupingPolicies(rules)
		}
	}

	rest.Success(c, "操作成功！", user)
}
