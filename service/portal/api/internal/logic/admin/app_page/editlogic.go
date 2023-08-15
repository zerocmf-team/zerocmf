package app_page

import (
	"context"
	"net/http"
	"strings"
	"zerocmf/service/portal/model"

	"gorm.io/gorm"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewEditLogic(header *http.Request, svcCtx *svc.ServiceContext) *EditLogic {
	ctx := header.Context()
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.AppPageSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id
	appPage := new(model.AppPage)
	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{0}
	query = append(query, "id = ?")
	queryArgs = append(queryArgs, id)
	queryStr := strings.Join(query, " AND ")
	err := appPage.Show(db, queryStr, queryArgs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.Error("该页面不存在或已被删除", err.Error())
			return
		}
		resp.Error("系统错误", err.Error())
		return
	}
	resp = savePage(db, req, 1)
	return
}
