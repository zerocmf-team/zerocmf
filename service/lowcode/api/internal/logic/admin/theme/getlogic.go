package theme

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

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

func (l *GetLogic) Get(req *types.ThemeListReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	current := 1
	pageSize := 10

	if req.Current != nil {
		current = *req.Current
	}

	if req.PageSize != nil {
		pageSize = *req.PageSize
	}

	var result interface{}
	result, err = new(model.Theme).List(db, current, pageSize, bson.M{})
	if err != nil {
		resp.Error("查询失败", err.Error())
	}

	resp.Success("获取成功！", result)
	return
}
