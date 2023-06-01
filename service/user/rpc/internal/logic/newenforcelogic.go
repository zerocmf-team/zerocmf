package logic

import (
	"context"
	"github.com/casbin/casbin/v2"
	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewEnforceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNewEnforceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewEnforceLogic {
	return &NewEnforceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NewEnforceLogic) NewEnforce(in *user.NewEnforceRequest) (reply *user.NewEnforcerReply, err error) {
	reply = new(user.NewEnforcerReply)
	c := l.svcCtx
	dbConf := c.Config.Database.NewConf(in.TenantId)
	var e *casbin.Enforcer
	e, err = dbConf.NewEnforcer()
	//	存入casbin
	if err != nil {
		return
	}

	userId := in.UserId
	menus := in.Menus
	var menusResult = make([]*user.Menu, 0)
	for _, v := range menus {
		path := v.Path
		if v.ParentId == 0 {
			path = v.Path
		}
		var access bool
		access, err = e.Enforce(userId, path, "*")
		if err != nil {
			return
		}

		if access {
			menusResult = append(menusResult, v)
		}
	}

	reply.Menus = menusResult
	return
}
