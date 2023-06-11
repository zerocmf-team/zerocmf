package category

import (
	"context"
	"strings"
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

func (l *GetLogic) Get(req *types.CateGetReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	name := req.Name
	if name != "" {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	status := req.Status
	if status != nil {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, status)
	}

	queryStr := strings.Join(query, " AND ")

	data, err := new(model.PortalCategory).Index(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return

}
