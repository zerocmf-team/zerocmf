package comment

import (
	"context"
	"net/http"
	"strconv"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyLikeLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewReplyLikeLogic(header *http.Request, svcCtx *svc.ServiceContext) *ReplyLikeLogic {
	ctx := header.Context()
	return &ReplyLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ReplyLikeLogic) ReplyLike(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}

	var query = "id = ? AND status = 1 AND delete_at = 0"
	var queryArgs = []interface{}{id}
	commentReply := new(model.CommentReply)
	err := commentReply.Show(db, query, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	if commentReply.Id == 0 {
		resp.Error("该回复不存在或已被删除", nil)
		return
	}
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{commentReply.Id, userId}

	postLikePost := model.CommentReplyLikePost{}

	prefix := c.Config.Database.Prefix
	postLikePost.Table = prefix + "comment_reply_like_post"

	postLike, status, err := postLikePost.Like(db, model.Post{
		Id:       commentReply.Id,
		UserId:   userIdInt,
		PostLike: commentReply.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"
	commentReply.IsLike = 1

	if status == false {
		msg = "取消点赞！"
		commentReply.IsLike = 0
	}

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", commentReply.Id).Model(&model.CommentReply{}).Update("post_like", postLike)

	commentReply.PostLike = postLike

	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success(msg, commentReply)
	return
}
