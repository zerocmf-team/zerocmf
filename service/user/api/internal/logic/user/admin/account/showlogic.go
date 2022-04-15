package account

import (
	"context"
	"gincmf/common/bootstrap/casbin"
	"gincmf/common/bootstrap/util"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"gincmf/service/user/model"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.OneReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	c := l.svcCtx
	db := c.Db

	resp = new(types.Response)
	id := req.Id
	if id == "" {
		resp.Error("id不能为空！", nil)
		return
	}

	user := model.User{}

	tx := db.Where("id = ? AND user_status = 1", []interface{}{id}).First(&user)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	if user.Id == 0 {
		resp.Error("该用户不存在！", nil)
		return
	}

	userId := strconv.Itoa(user.Id)

	//	获取当前用户的全部角色
	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	roles, err := e.GetRolesForUser(userId)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var result struct {
		model.User
		Roles []string `json:"roles"`
	}

	result.User = user
	result.Roles = roles

	resp.Success("获取成功！", result)


	return
}
