package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"zerocmf/common/bootstrap/database"

	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatabaseLogic {
	return &DatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DatabaseLogic) Database(in *user.DatabaseRequest) (*user.DatabaseReply, error) {
	dbConf := database.Conf()
	reply := new(user.DatabaseReply)
	err := copier.Copy(&reply, &dbConf)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
