package post

import (
	"context"
	"net/http"
	"strconv"
	comModel "zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(header *http.Request, svcCtx *svc.ServiceContext) *LikeLogic {
	ctx := header.Context()
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}

	var query = "id = ? AND post_type = ? and delete_at = ?"
	var queryArgs = []interface{}{id, 1, 0}

	post := new(model.PortalPost)

	err := post.Show(db, query, queryArgs)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	if post.Id == 0 {
		resp.Error("该文章不存在或已被删除", nil)
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{post.Id, userId}

	postLikePost := model.PostLikePost{}

	prefix := c.Config.Database.Prefix
	postLikePost.Table = prefix + "post_like_post"

	postLike, status, err := postLikePost.Like(db, comModel.Post{
		Id:       post.Id,
		UserId:   userIdInt,
		PostLike: post.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"

	if status == false {
		msg = "取消点赞！"
	}

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", post.Id).Model(&model.PortalPost{}).Update("post_like", postLike)

	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success(msg, post)
	return
}
