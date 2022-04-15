package role

import (
	"context"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"gincmf/service/user/model"
	"strings"

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
	resp = new(types.Response)
	c := l.svcCtx
	db := c.Db

	id := req.Id
	if id == "" {
		resp.Error("角色id不能为空！", nil)
		return
	}

	role := model.Role{}
	query := []string{"id = ?", "status = 1"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{id}
	err = role.Show(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", role)
	return
}
