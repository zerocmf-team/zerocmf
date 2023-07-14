package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	cmfRedis "zerocmf/common/bootstrap/redis"
	"zerocmf/service/admin/rpc/internal/svc"
	"zerocmf/service/admin/rpc/types/admin"
)

type EncryptUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEncryptUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EncryptUidLogic {
	return &EncryptUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EncryptUidLogic) EncryptUid(in *admin.EncryptUidReq) (reply *admin.EncryptUidReply, err error) {
	reply = new(admin.EncryptUidReply)
	c := l.svcCtx
	redis := c.Redis

	key := in.GetKey()
	salt := in.GetSalt()

	var uid int64
	uid, err = cmfRedis.NewRedis(redis).EncryptUid(key, int(salt))
	if err != nil {
		return
	}
	reply.Uid = uid
	return
}
