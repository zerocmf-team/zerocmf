package navItem

import (
	"context"
	"strconv"
	"strings"
	"zerocmf/service/portal/model"

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

type TreeOption struct {
	Label    string       `json:"label"`
	Value    string       `json:"value"`
	Children []TreeOption `json:"children"`
}

func recursionTreeOption(data []model.PortalTree) (option []TreeOption) {
	for _, item := range data {
		var url = "/list/" + strconv.Itoa(item.Id)
		if item.Alias != "" {
			url = "/" + item.Alias
		}
		// 单个选项
		treeSelect := TreeOption{
			Label:    item.Name,
			Value:    url,
			Children: make([]TreeOption, 0),
		}
		if len(item.Children) > 0 {
			children := recursionTreeOption(item.Children)
			treeSelect.Children = children
		}
		option = append(option, treeSelect)
	}
	return
}

func (l *OptionsUrlsLogic) OptionsUrls() (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{"0"}

	queryStr := strings.Join(query, " AND ")

	categoryData, err := new(model.PortalCategory).Index(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	cateOptions := recursionTreeOption(categoryData)

	tree := make([]TreeOption, 0)
	tree = append(tree, TreeOption{
		Label:    "首页",
		Value:    "/",
		Children: make([]TreeOption, 0),
	})

	postQuery := "post_type = ?"
	postQueryArgs := []interface{}{"2"}
	post := new(model.PortalPost)
	pages, pageErr := post.PortalList(db, postQuery, postQueryArgs)
	if pageErr != nil {
		resp.Error(pageErr.Error(), nil)
		return
	}

	for _, v := range cateOptions {
		tree = append(tree, TreeOption{
			Label:    v.Label,
			Value:    v.Value,
			Children: v.Children,
		})
	}

	for _, v := range pages {
		value := "/page/" + strconv.Itoa(v.Id)
		if v.MoreJson.Alias != "" {
			value = v.MoreJson.Alias
		}
		tree = append(tree, TreeOption{
			Label:    v.PostTitle,
			Value:    value,
			Children: make([]TreeOption, 0),
		})
	}
	resp.Success("获取成功！", tree)
	return
}
