package account

import (
	"context"
	"zerocmf/common/bootstrap/casbin"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"
	"strconv"
	"time"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.AdminSaveReq) (resp types.Response) {

	editId := req.Id
	userId, _ := strconv.Atoi(req.Id)

	c := l.svcCtx
	db := c.Db

	form := req

	if len(form.RoleIds) <= 0 {
		resp.Error("至少选择一项角色！", nil)
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

	// 存入用户角色
	e, err := casbin.NewEnforcer("")

	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	editUser := model.User{}
	tx := db.Where("id = ?", editId).First(&editUser)

	if editUser.Id > 0 && editUser.UserLogin != form.UserLogin {
		currentUser := model.User{}
		tx = db.Where("user_login = ?", form.UserLogin).First(&currentUser)
		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}
		if currentUser.Id > 0 {
			resp.Error("该登录名已存在！", nil)
			return
		}
	}

	if form.UserPass == "" {
		user.UserPass = util.GetMd5(form.UserPass)
	}
	user.Id = userId
	tx = db.Save(&user)

	if tx.Error != nil {
		resp.Error("创建用户出错，请联系管理员！"+tx.Error.Error(), nil)
		return
	}

	roles, err := e.GetRolesForUser(editId)
	if err != nil {
		resp.Error(err.Error(), nil)
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
		rules = append(rules, []string{editId, v})
	}
	if len(rules) > 0 {
		e.RemoveGroupingPolicies(rules)
	}

	// 开始新增策略
	rules = make([][]string, 0)
	for _, v := range alreadyAdd {
		rules = append(rules, []string{editId, v})
	}
	if len(rules) > 0 {
		e.AddGroupingPolicies(rules)
	}

	resp.Success("更新成功！", user)
	return
}
