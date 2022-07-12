package navItem

import (
	"context"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionsListLogic {
	return &OptionsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionsListLogic) OptionsList(req *types.NavItemOptionsReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	navId := req.NavId
	if navId == 0 {
		resp.Error("导航不能为空！", nil)
		return
	}

	var query = "nav_id  = ?"
	var queryArgs = []interface{}{navId}

	result := new(model.NavItem).OptionsList(db, query, queryArgs)
	resp.Success("获取成功！", result)
	return

}
