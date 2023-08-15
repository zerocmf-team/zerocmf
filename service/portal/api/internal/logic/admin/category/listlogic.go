package category

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewListLogic(header *http.Request, svcCtx *svc.ServiceContext) *ListLogic {
	ctx := header.Context()
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	Category := model.PortalCategory{
		ParentId: 0,
	}
	data, err := Category.ListWithTree(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
