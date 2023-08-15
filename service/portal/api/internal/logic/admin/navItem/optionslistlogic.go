package navItem

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsListLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewOptionsListLogic(header *http.Request, svcCtx *svc.ServiceContext) *OptionsListLogic {
	ctx := header.Context()
	return &OptionsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *OptionsListLogic) OptionsList(req *types.NavItemOptionsReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

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
