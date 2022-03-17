package logic

import (
	"context"
	"fmt"
	"gincmf/service/user/model"
	"gincmf/service/user/rpc/internal/svc"
	"gincmf/service/user/rpc/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitLogic {
	return &InitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitLogic) Init(in *userclient.InitRequest) (*userclient.InitReply, error) {
	// todo: add your logic here and delete this line

	fmt.Println("Timeout",l.svcCtx.Config.Timeout)
	tenantId := in.GetTenantId()
	fmt.Println("tenantId",tenantId)
	model.Migrate(tenantId)

	time.Sleep(time.Second * 5)

	fmt.Println("同步完成123 ok")

	return &userclient.InitReply{
		Code: 1,
		Message: "数据库迁移同步完成",
	}, nil
}
