package category

import (
	"context"
	"net/http"
	"strings"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewOptionsLogic(header *http.Request, svcCtx *svc.ServiceContext) *OptionsLogic {
	ctx := header.Context()
	return &OptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *OptionsLogic) Options() (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	Category := model.PortalCategory{}
	var query = []string{"delete_at  = ?"}
	var queryArgs = []interface{}{0}
	queryStr := strings.Join(query, " AND ")

	data, err := Category.ListWithOptions(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
