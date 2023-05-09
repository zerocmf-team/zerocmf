package nav

import (
	"context"
	"strings"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

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

func (l *ShowLogic) Show(req *types.NavShowReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	var nav = new(model.Nav)
	var query = []string{"id = ?"}
	var queryArgs = []interface{}{req.Id}
	queryStr := strings.Join(query, " AND ")
	err := nav.Show(db, queryStr, queryArgs)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}
	if nav.Id == 0 {
		resp.Error("该导航不存在或已被删除！", nil)
		return
	}
	resp.Success("获取成功！", nav)
	return
}
