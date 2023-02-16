package post

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.PostShowReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}

	typ := req.Type
	if !(typ == 1 || typ == 2) {
		resp.Error("页面类型错误", nil)
		return
	}

	var query = "id = ? AND post_type = ? and delete_at = ?"
	var queryArgs = []interface{}{id, typ, 0}

	post := model.PortalPost{}

	err := post.Show(db, query, queryArgs)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	if post.Id == 0 {
		resp.Error("该文章不存在或已被删除", nil)
		return
	}

	// 更新访问量 +1
	postHits := post.PostHits
	postHits += 1
	post.PostHits = postHits

	tx := db.Model(model.PortalPost{Id: post.Id}).Update("post_hits", postHits)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	var result struct {
		model.PortalPost
		PrevPost *model.PortalPost `json:"prev_post"`
		NextPost *model.PortalPost `json:"next_post"`
	}

	result.PortalPost = post

	if typ == 1 {

		// 查询上一篇
		query = "id < ? AND post_type = ? and delete_at = ?"
		queryArgs = []interface{}{id, typ, 0}

		prevPost := model.PortalPost{}
		err = prevPost.Show(db, query, queryArgs)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		// 查询下一篇
		query = "id > ? AND post_type = ? and delete_at = ?"
		queryArgs = []interface{}{id, 1, 0}
		nextPost := model.PortalPost{}

		err = nextPost.Show(db, query, queryArgs)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		if prevPost.Id > 0 {
			result.PrevPost = &prevPost
		}

		if nextPost.Id > 0 {
			result.NextPost = &nextPost
		}
	}

	resp.Success("获取成功！", result)
	return
}
