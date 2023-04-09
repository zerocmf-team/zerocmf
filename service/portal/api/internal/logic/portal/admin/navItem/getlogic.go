package navItem

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.NavItemGetReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	r := c.Request

	query := "id = ?"
	queryArgs := []interface{}{req.NavId}
	nav := model.Nav{}
	err := nav.Show(db, query, queryArgs)
	if err != nil {
		resp.Error("操作失败", err.Error())
		return
	}

	// 根据navId获取全部导航项
	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	itemQuery := "nav_id = ? AND status = 1"
	itemQueryArgs := []interface{}{req.NavId}

	navItemsPaginate, err := new(model.NavItem).GetWithChildPaginate(db, current, pageSize, itemQuery, itemQueryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	navItems, err := new(model.NavItem).GetWithChild(db, itemQuery, itemQueryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("操作成功！", gin.H{
		"navId":            nav.Id,
		"navItemsPaginate": navItemsPaginate,
		"navItems":         navItems,
	})
	return
}
