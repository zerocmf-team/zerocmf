package navItem

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.NavItemGetReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := l.header

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
