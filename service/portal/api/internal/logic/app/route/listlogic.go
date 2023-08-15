package route

import (
	"context"
	"net/http"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

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
	data, err := new(model.Route).List(db, "", nil)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
