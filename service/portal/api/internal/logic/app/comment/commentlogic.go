package comment

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/user/rpc/types/user"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(header *http.Request, svcCtx *svc.ServiceContext) *CommentLogic {
	ctx := header.Context()
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.PostCommentAddReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	userRpc := c.UserRpc

	topicId := req.Id
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	var (
		userData *user.UserReply
		err      error
	)

	if userId != "" {

		userData, err = userRpc.Get(context.Background(), &user.UserRequest{
			UserId: userId.(string),
			SiteId: siteId.(int64),
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

	}

	now := time.Now().Unix()
	comment := model.Comment{
		TopicId:          topicId,
		TopicType:        req.TopicType,
		Content:          req.Content,
		FromUserId:       userIdInt,
		FromUserNickname: userData.UserNickname,
		FromUserAvatar:   userData.Avatar,
	}

	comment.CreateAt = now
	comment.UpdateAt = now

	tx := db.Create(&comment)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success("评论成功！", comment)
	return
}
