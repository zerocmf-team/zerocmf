package tag

import (
	"context"
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

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

func (l *GetLogic) Get(req *types.TagGetReq) (resp types.Response) {

	c := l.svcCtx
	r := l.header
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	name := req.Name

	var query []string
	var queryArgs []interface{}

	if name != "" {
		query = []string{"name like ?"}
		queryArgs = []interface{}{"%" + name + "%"}
	}

	current, pageSize, err := data.NewPaginate(r).Default()

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	queryStr := strings.Join(query, " AND ")
	data, err := new(model.PortalTag).Index(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
