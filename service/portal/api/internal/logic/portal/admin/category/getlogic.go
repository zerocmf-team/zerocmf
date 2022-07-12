package category

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"
	"strings"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

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

func (l *GetLogic) Get(req *types.CateGetReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	db := c.Db

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	name := req.Name
	if name != "" {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := new(data.Paginate).Default(r)

	data, err := new(model.PortalCategory).Index(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return

}
