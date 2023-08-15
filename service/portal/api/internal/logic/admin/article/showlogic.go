package article

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	comModel "zerocmf/common/bootstrap/model"
	"zerocmf/common/bootstrap/util"
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

func (l *ShowLogic) Show(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	gormDB := c.NewDb(siteId.(int64))
	id := req.Id

	query := []string{"id = ?", "delete_at = ?"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{id, 0}

	var result struct {
		model.PortalPost
		UserLogin string                 `json:"user_login"`
		Alias     string                 `json:"alias"`
		Keywords  []string               `json:"keywords"`
		Category  []model.PortalCategory `json:"Category"`
		Extends   []model.Extends        `json:"extends"`
		Slug      string                 `json:"slug"`
		model.More
	}

	post := model.PortalPost{}

	// 获取当前文章信息
	err := post.Show(gormDB.Db, queryStr, queryArgs)
	if err != nil {
		resp.Error("查询失败："+err.Error(), nil)
		return
	}

	result.PortalPost = post

	if post.PostKeywords != "" {
		result.Keywords = strings.Split(post.PostKeywords, ",")
	}

	// 获取当前文章全部分类
	pQueryArgs := []interface{}{id, 0}
	pCate := model.PortalCategory{}
	Category, err := pCate.FindPostCategory(gormDB, "p.id = ? AND p.delete_at = ?", pQueryArgs)
	result.Category = Category

	result.Extends = post.MoreJson.Extends
	result.ExtendsObj = post.MoreJson.ExtendsObj

	result.Photos = post.MoreJson.Photos
	result.Files = post.MoreJson.Files

	result.Audio = post.MoreJson.Audio
	result.AudioPrevPath = post.MoreJson.AudioPrevPath

	result.Video = post.MoreJson.Video
	result.VideoPrevPath = post.MoreJson.VideoPrevPath

	fullUrl := "page/" + strconv.Itoa(id)
	route := comModel.Route{}
	tx := gormDB.Db.Where("full_url", fullUrl).First(&route)

	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	result.Alias = route.Url
	resp.Success("获取成功！", result)
	return
}
