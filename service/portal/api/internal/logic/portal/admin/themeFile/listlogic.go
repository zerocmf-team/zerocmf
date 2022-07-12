package themeFile

import (
	"context"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

type option struct {
	Id   int    `json:"id"`
	File string `json:"file"`
	Name string `json:"name"`
}

func (l *ListLogic) List(req *types.ListReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	t := req.Type
	opt := model.Option{}

	tx := db.Where("option_name = ?", "theme").First(&opt)
	if tx.Error != nil {
		resp.Error("该主题不存在！", nil)
		if tx.Error != gorm.ErrRecordNotFound {
			resp.Error(tx.Error.Error(), nil)
			return
		}
		return
	}

	theme := opt.OptionValue
	var list []model.ThemeFile
	tx = db.Where("theme = ? AND type = ?", theme, t).Find(&list)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	file := "list"
	if t == "article" {
		file = "article"
	} else if t == "page" {
		file = "page"
	}

	result := []option{{Name: "默认模板", File: file}}
	for _, v := range list {
		result = append(result, option{
			Id:   v.Id,
			Name: v.Name,
			File: v.File,
		})
	}

	resp.Success("获取成功！", result)

	return
}
