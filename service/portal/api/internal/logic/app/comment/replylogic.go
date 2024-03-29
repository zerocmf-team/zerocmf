package comment

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/user/rpc/types/user"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewReplyLogic(header *http.Request, svcCtx *svc.ServiceContext) *ReplyLogic {
	ctx := header.Context()
	return &ReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ReplyLogic) Reply(req *types.PostReplyReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	userRpc := c.UserRpc

	id := req.Id

	userId, _ := c.Get("userId")

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

	toUserId := req.ToUserId

	now := time.Now().Unix()
	reply := model.CommentReply{}
	copier.Copy(&reply, &req)

	reply.CommentId = id
	reply.FromUserId = int(userData.Id)
	reply.FromUserNickname = userData.UserNickname
	reply.FromUserAvatar = userData.Avatar
	reply.CreateAt = now
	reply.UpdateAt = now

	var (
		toUserData *user.UserReply
	)

	if toUserId != 0 {

		userData, err = userRpc.Get(context.Background(), &user.UserRequest{
			UserId: userId.(string),
			SiteId: siteId.(int64),
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		reply.ToUserId = int(toUserData.Id)
		reply.ToUserNickname = toUserData.UserNickname
		reply.ToUserAvatar = toUserData.Avatar
	}

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	tx := db.Create(&reply)
	err = tx.Error
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("回复成功！", reply)
	return
}
