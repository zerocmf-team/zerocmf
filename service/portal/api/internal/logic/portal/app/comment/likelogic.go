package comment

import (
	"context"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}

	var query = "id = ? AND status = 1 AND delete_at = 0"
	var queryArgs = []interface{}{id}

	comment := new(model.Comment)
	err := comment.Show(db, query, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	if comment.Id == 0 {
		resp.Error("该评论不存在或已被删除", nil)
		return
	}

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	query = "post_id = ? AND user_id = ?"
	queryArgs = []interface{}{comment.Id, userId}

	postLikePost := model.CommentLikePost{}
	prefix := c.Config.Database.Prefix
	postLikePost.Table = prefix + "comment_like_post"

	postLike, status, err := postLikePost.Like(db, model.Post{
		Id:       comment.Id,
		UserId:   userIdInt,
		PostLike: comment.PostLike,
	}, query, queryArgs)

	msg := "点赞成功！"
	comment.IsLike = 1

	if status == false {
		msg = "取消点赞！"
		comment.IsLike = 0
	}

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	// 更新post_like
	tx := db.Where("id", comment.Id).Model(&model.Comment{}).Update("post_like", postLike)
	comment.PostLike = postLike
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success(msg, comment)
	return
}
