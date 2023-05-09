package app_page

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.AppPageSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
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
