package themeFile

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(header *http.Request, svcCtx *svc.ServiceContext) *DetailLogic {
	ctx := header.Context()
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.ThemeFileDetailReq) (resp *types.Response) {
	c := l.svcCtx
	theme := req.Theme

	resp = new(types.Response)

	if theme == "" {
		resp.Error("主题不能为空！", nil)
		return
	}

	file := req.File
	if file == "" {
		resp.Error("文件不能为空！", nil)
		return
	}

	query := "theme = ? AND file = ?"
	queryArgs := []interface{}{theme, file}

	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	data, err := new(model.ThemeFile).Show(db, query, queryArgs)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
