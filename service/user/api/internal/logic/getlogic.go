package logic

import (
	"context"
	"fmt"
	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"
	"gincmf/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{}
	c := l.svcCtx
	uRpc := c.UserRpc

	fmt.Println("hit Get Get Get",c.Config.Timeout)

	res, err := uRpc.Init(l.ctx, &user.InitRequest{
		TenantId: "123456",
	})

	resp.Error(res.GetMessage(), nil)
	if err != nil {
		return
	}

	resp.Error(res.GetMessage(), nil)
	if res.GetCode() == 0 {
		return
	}

	resp.Success(res.GetMessage(), nil)
	return
}
