package account

import (
	"context"
	"gincmf/common/bootstrap/casbin"
	"gincmf/common/bootstrap/util"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"gincmf/service/user/model"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.AdminStoreReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	c := l.svcCtx
	db := c.Db

	resp = new(types.Response)
	form := req
	if len(form.RoleIds) <= 0 {
		resp.Error("至少选择一项角色！", nil)
		return
	}

	user := model.User {
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

	currentUser := model.User{}
	tx := db.Where("user_login = ?", form.UserLogin).First(&currentUser)

	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if currentUser.Id > 0 {
		resp.Error("该用户已存在！", nil)
		return
	}

	tx = db.Create(&user)
	if tx.Error != nil {
		resp.Error("创建用户出错，请联系管理员！", tx.Error)
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

	resp.Success("操作成功！", user)

	return
}
