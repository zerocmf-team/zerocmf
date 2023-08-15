package form

import (
	"context"
	"net/http"
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
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.FormListReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := l.header

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{0}

	if req.Name != nil {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+*req.Name+"%")
	}

	if req.Status != nil {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, req.Status)
	}

	queryStr := strings.Join(query, " AND ")
	form := model.Form{}
	res, resErr := form.List(db, current, pageSize, queryStr, queryArgs, true)
	if resErr != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功!", res)
	return
}
