package form

import (
	"context"
	"net/http"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewShowLogic(header *http.Request, svcCtx *svc.ServiceContext) *ShowLogic {
	ctx := header.Context()
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.FormShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	form := model.Form{}
	queryArgs := []interface{}{req.Id, 1}
	err := form.Show(db, "id = ? and status = ?", queryArgs)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}
	resp.Success("获取成功!", form)
	return
}
