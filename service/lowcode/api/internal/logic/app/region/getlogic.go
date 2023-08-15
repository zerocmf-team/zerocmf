package region

import (
	"context"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.RegionGetReq) (resp types.Response) {
	c := l.svcCtx
	// 选择租户表
	db, err := c.MongoDB()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	var list []model.Region
	list, err = new(model.Region).List(db)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}
	resp.Success("获取成功", list)
	return
}
