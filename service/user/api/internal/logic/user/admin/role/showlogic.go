package role

import (
	"context"
	"strings"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"zerocmf/service/user/model"

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

func (l *ShowLogic) Show(req *types.OneReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))

	id := req.Id
	if id == "" {
		resp.Error("角色id不能为空！", nil)
		return
	}

	role := model.Role{}
	query := []string{"id = ?", "status = 1"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{id}
	err := role.Show(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", role)
	return
}
