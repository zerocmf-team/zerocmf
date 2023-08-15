package nav

import (
	"context"
	"net/http"
	"strings"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(header *http.Request, svcCtx *svc.ServiceContext) *DeleteLogic {
	ctx := header.Context()
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.NavShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	var nav = new(model.Nav)
	var query = []string{"id = ?"}
	var queryArgs = []interface{}{req.Id}
	queryStr := strings.Join(query, " AND ")
	err := nav.Show(db, queryStr, queryArgs)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}
	if nav.Id == 0 {
		resp.Error("该导航不存在或已被删除！", nil)
		return
	}
	tx := db.Where("id = ?", nav.Id).Delete(&nav)
	if tx.Error != nil {
		resp.Error("系统错误", err.Error())
		return
	}
	resp.Success("删除成功！", nav)
	return
}
