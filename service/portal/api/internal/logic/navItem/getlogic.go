package navItem

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
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
	key := req.Key

	if key == "" {
		resp.Error("唯一标识不能为空！", nil)
		return
	}

	query := "`key` = ?"
	queryArgs := []interface{}{key}
	nav := model.Nav{}
	err := nav.Show(db, query, queryArgs)
	if err != nil {
		resp.Error("操作失败", nil)
		return
	}

	if nav.Id == 0 {
		nav.Key = key
		tx := db.Where(query, queryArgs...).FirstOrCreate(&nav)
		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}
	}

	// 根据navId获取全部导航项
	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error( err.Error(), nil)
		return
	}

	itemQuery := "nav_id = ? AND status = 1"
	itemQueryArgs := []interface{}{nav.Id}

	navItemsPaginate, err := new(model.NavItem).GetWithChildPaginate(db, current, pageSize, itemQuery, itemQueryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	navItems, err := new(model.NavItem).GetWithChild(db, itemQuery, itemQueryArgs)
	if err != nil {
		resp.Error( err.Error(), nil)
		return
	}

	resp.Success("操作成功！", gin.H{
		"navId":            nav.Id,
		"navItemsPaginate": navItemsPaginate,
		"navItems":         navItems,
	})
	return
}
