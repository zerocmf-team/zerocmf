package themeFile

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ThemeFileListReq) (resp types.Response) {
	c := l.svcCtx
	theme := req.Theme
	if theme == "" {
		resp.Error("主题不能为空！", nil)
		return
	}

	isPublic := req.IsPublic

	query := "theme = ? AND is_public = ?"
	queryArgs := []interface{}{theme, isPublic}

	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))

	data, err := new(model.ThemeFile).List(db, query, queryArgs)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
