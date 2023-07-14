package account

import (
	"context"
	"strconv"
	"strings"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FirstOrCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFirstOrCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FirstOrCreateLogic {
	return &FirstOrCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FirstOrCreateLogic) FirstOrCreate(req *types.AdminStoreReq) (resp types.Response) {
	c := l.svcCtx

	siteId, _ := c.Get("siteId")
	dbConf := c.Config.Database.NewConf(siteId.(string))
	db := dbConf.ManualDb(siteId.(string))

	form := req
	if len(form.RoleIds) <= 0 {
		resp.Error("至少选择一项角色！", nil)
		return
	}

	var query = []string{"delete_at = ?"}
	var queryArgs = []interface{}{0}

	if form.UserLogin == "" && form.Mobile == "" {
		resp.Error("登录账号或手机不能为空！", nil)
		return
	}

	if form.UserLogin != "" {
		query = append(query, "user_login = ?")
		queryArgs = append(queryArgs, form.UserLogin)
	}

	if form.Mobile != "" {
		query = append(query, "mobile = ?")
		queryArgs = append(queryArgs, form.Mobile)
	}

	queryStr := strings.Join(query, " AND ")

	existUser := model.User{}
	tx := db.Where(queryStr, queryArgs...).First(&existUser)

	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if existUser.Id > 0 {
		resp.Success("操作成功！", existUser)
		return
	}

	userPass := util.GetMd5(form.UserPass)

	user := model.User{
		UserType:     1,
		CreateAt:     time.Now().Unix(),
		Mobile:       form.Mobile,
		UserRealName: form.UserRealname,
		UserLogin:    form.UserLogin,
		UserPass:     userPass,
		UserEmail:    form.UserEmail,
		UserStatus:   1,
	}

	// 存入用户角色
	e, err := dbConf.NewEnforcer()
	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	tx = db.Create(&user)
	if tx.Error != nil {
		resp.Error("创建用户出错，请联系管理员！", tx.Error)
		return
	}

	userId := strconv.Itoa(user.Id)
	roleIds := form.RoleIds
	rules := make([][]string, 0)
	for _, v := range roleIds {
		rules = append(rules, []string{userId, strconv.Itoa(v)})
	}
	if len(rules) > 0 {
		e.AddGroupingPolicies(rules)
	}

	resp.Success("操作成功！", user)
	return
}
