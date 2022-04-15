package tag

import (
	"context"
	"gincmf/common/bootstrap/data"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"
	"strings"

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

func (l *GetLogic) Get(req *types.TagGetReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	db := c.Db

	name := req.Name

	var query []string
	var queryArgs []interface{}

	if name != "" {
		query = []string{"name like ?"}
		queryArgs = []interface{}{"%" + name + "%"}
	}

	current, pageSize, err := new(data.Paginate).Default(r)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	queryStr := strings.Join(query, " AND ")
	data, err := new(model.PortalTag).Index(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		resp.Error( err.Error(), nil)
		return
	}

	resp.Success( "获取成功！", data)
	return
}
