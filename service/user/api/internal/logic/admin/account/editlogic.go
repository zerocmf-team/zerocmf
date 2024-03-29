package account

import (
	"context"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"

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
	siteId, _ := c.Get("siteId")

	dbConf := c.Config.Database.NewConf(siteId.(int64))
	db := dbConf.ManualDb(siteId.(int64))

	if len(req.RoleIds) <= 0 {
		resp.Error("至少选择一项角色！", nil)
		return
	}

	user := model.User{
		UserType:     1,
		CreateAt:     time.Now().Unix(),
		Mobile:       req.Mobile,
		UserRealName: req.UserRealname,
		UserLogin:    req.UserLogin,
		UserEmail:    req.UserEmail,
		UserStatus:   1,
	}

	// 存入用户角色
	e, err := dbConf.NewEnforcer()

	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	editUser := model.User{}
	tx := db.Where("id = ?", editId).First(&editUser)

	// 查询用户名是否唯一
	if editUser.Id > 0 && editUser.UserLogin != req.UserLogin {
		existUser := model.User{}
		tx = db.Where("user_login = ?", req.UserLogin).First(&existUser)
		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}
		if existUser.Id > 0 {
			resp.Error("该登录名已存在！", nil)
			return
		}
	}

	if req.UserPass == "" {
		user.UserPass = editUser.UserPass
	} else {
		user.UserPass = util.GetMd5(req.UserPass)
	}
	user.Id = userId
	tx = db.Save(&user)

	if tx.Error != nil {
		resp.Error("创建用户出错，请联系管理员！"+tx.Error.Error(), nil)
		return
	}

	var roles []string
	roles, err = e.GetRolesForUser(editId)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	alreadyDel := make([]string, 0)
	// 判断是否需要被删除
	for _, v := range roles {
		if util.ToLowerInArray(v, req.RoleIds) == false {
			alreadyDel = append(alreadyDel, v)
		}
	}
	// 如果新增为空，则全部删除
	if len(req.RoleIds) == 0 {
		alreadyDel = roles
	}

	alreadyAdd := make([]string, 0)
	for _, v := range req.RoleIds {
		roleId := strconv.Itoa(v)
		if util.ToLowerInArray(roleId, roles) == false {
			alreadyAdd = append(alreadyAdd, roleId)
		}
	}

	// 如果数据库不存在，则为新增
	if len(roles) == 0 {
		alreadyAdd = make([]string, 0)
		for _, v := range req.RoleIds {
			alreadyAdd = append(alreadyAdd, strconv.Itoa(v))
		}
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
