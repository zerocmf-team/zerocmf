package comment

import (
	"context"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/user/rpc/user"
	"strconv"
	"time"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.PostCommentAddReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	userRpc := c.UserRpc

	topicId := req.Id
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	var (
		userData *user.UserReply
		err error
	)

	if userId != "" {
		tenant, exist := db.Get("tenantId")
		tenantId := ""
		if exist {
			tenantId = tenant.(string)
		}

		userData, err = userRpc.Get(context.Background(), &user.UserRequest{
			UserId:   int64(userIdInt),
			TenantId: tenantId,
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
