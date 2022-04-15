package themeFile

import (
	"context"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.ThemeFileDetailReq) (resp types.Response) {
	c := l.svcCtx
	theme := req.Theme

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

	db := c.Db

	data, err := new(model.ThemeFile).Show(db, query, queryArgs)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
