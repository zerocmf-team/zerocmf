package list

import (
	"context"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"

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

func (l *GetLogic) Get(req *types.PostListReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	r := c.Request

	extra := map[string]string{}

	hot := req.Hot

	if hot == 1 {
		extra["hot"] = "1"
	}

	var query []string
	var queryArgs []interface{}

	ids := req.Ids

	idsArr := strings.Split(ids, ",")
	for _, v := range idsArr {
		query = append(query, "cp.category_id = ?")
		queryArgs = append(queryArgs, v)
	}

	queryRes := []string{"p.post_type = 1 AND p.delete_at = 0"}
	if len(query) > 0 {
		queryStr := strings.Join(query, " OR ")
		queryRes = append(queryRes, queryStr)
	}

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	data, err := new(model.PortalPost).ListByCategory(db, current, pageSize, strings.Join(queryRes, " AND "), queryArgs, extra)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
