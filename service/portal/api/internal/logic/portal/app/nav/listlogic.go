package nav

import (
	"context"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.NavItemGetReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db

	query := "id = ?"
	queryArgs := []interface{}{req.NavId}
	nav := model.Nav{}
	err := nav.Show(db, query, queryArgs)
	if err != nil {
		resp.Error("操作失败", err.Error())
		return
	}

	// 根据navId获取全部导航项
	itemQuery := "nav_id = ? AND status = 1"
	itemQueryArgs := []interface{}{req.NavId}

	navItems, err := new(model.NavItem).GetWithChild(db, itemQuery, itemQueryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("操作成功！", navItems)
	return
}
