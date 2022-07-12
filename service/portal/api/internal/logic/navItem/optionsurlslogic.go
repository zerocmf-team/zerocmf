package navItem

import (
	"context"
	"zerocmf/service/portal/model"
	"strconv"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsUrlsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionsUrlsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionsUrlsLogic {
	return &OptionsUrlsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Options struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionsMap struct {
	Label   string    `json:"label"`
	Options []Options `json:"options"`
}

func (l *OptionsUrlsLogic) OptionsUrls() (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	portalCategory, err := new(model.PortalCategory).List(db)

	if err != nil {
		resp.Error(err.Error(), nil)
	}

	categoryOptions := make([]Options, 0)

	for _, v := range portalCategory {

		var url = "/list/" + strconv.Itoa(v.Id)
		if v.Alias != "" {
			url = "/" + v.Alias
		}

		categoryOptions = append(categoryOptions, Options{
			Label: v.Name,
			Value: url,
		})

	}

	query := "post_type = ?"
	queryArgs := []interface{}{"2"}

	pages, err := model.PortalPost{}.PortalList(db, query, queryArgs)

	pageOptions := make([]Options, 0)

	for _, v := range pages {

		value := "/page/" + strconv.Itoa(v.Id)

		if v.MoreJson.Alias != "" {
			value = v.MoreJson.Alias
		}

		pageOptions = append(pageOptions, Options{
			Label: v.PostTitle,
			Value: value,
		})

	}

	var om = []OptionsMap{{
		Label: "首页",
		Options: []Options{
			{
				Label: "首页",
				Value: "/",
			}},
	}, {
		Label:   "文章分类",
		Options: categoryOptions,
	}, {
		Label:   "所有页面",
		Options: pageOptions,
	}}

	resp.Success("获取成功！", om)

	return
}
