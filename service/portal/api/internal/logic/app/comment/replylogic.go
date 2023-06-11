package comment

import (
	"context"
	"github.com/jinzhu/copier"
	"strconv"
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
	svcCtx *svc.ServiceContext
}

func NewReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyLogic {
	return &ReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyLogic) Reply(req *types.PostReplyReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	userRpc := c.UserRpc

	id := req.Id

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	var (
		userData *user.UserReply
		err      error
	)

	tenant, exist := db.Get("tenantId")
	tenantId := ""
	if exist {
		tenantId = tenant.(string)
	}

	if userId != "" {

		userData, err = userRpc.Get(context.Background(), &user.UserRequest{
			UserId:   int64(userIdInt),
			TenantId: tenantId,
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
			UserId:   int64(userIdInt),
			TenantId: tenantId,
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
