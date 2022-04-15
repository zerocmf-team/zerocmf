package logic

import (
	"context"
	"gincmf/common/bootstrap/data"
	"gincmf/service/admin/api/internal/svc"
	"gincmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)
	resp.Success("获取成功！", data.H{"version": "v1", "message": "hello admin v1"})

	return
}
