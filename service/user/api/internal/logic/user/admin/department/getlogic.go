package department

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"zerocmf/common/bootstrap/required"
	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"
	"zerocmf/service/user/model"
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

func (l *GetLogic) Get(req *types.DepListReq) (resp *types.Response) {
	resp = new(types.Response)
	c := l.svcCtx
	db := c.Db
	query := make([]string, 0)
	var queryArgs = make([]interface{}, 0)
	name := req.Name
	if required.String(name) {
		query = append(query, "name like %?%")
		queryArgs = append(queryArgs, name)
	}
	status := req.Status
	if required.String(status) {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, name)
	}

	queryStr := strings.Join(query, " AND ")

	department := model.Department{}
	res, err := department.TreeList(db, queryStr, queryArgs)
	if err != nil {
		resp.Error("系统错误", nil)
		return
	}
	resp.Success("操作成功", res)
	return
}
