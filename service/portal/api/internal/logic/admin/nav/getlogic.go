package nav

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

func (l *GetLogic) Get(req *types.NavGetReq) (resp types.Response) {

	c := l.svcCtx
	r := l.header
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	var nav = new(model.Nav)

	// 根据navId获取全部导航项
	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var query []string
	var queryArgs []interface{}
	if req.Name != "" {
		query = append(query, "name = ?")
		queryArgs = append(queryArgs, req.Name)
	}
	queryStr := strings.Join(query, " AND ")
	var data data.Paginate
	data, err = nav.Get(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		return types.Response{}
	}
	if err != nil {
		resp.Error("系统出错！", err.Error())
		return
	}
	resp.Success("获取成功!", data)
	return
}
