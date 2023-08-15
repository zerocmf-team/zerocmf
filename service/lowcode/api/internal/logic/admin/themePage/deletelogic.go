package themePage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.ThemePageShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("themePage")

	id := req.Id

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		resp.Error("非法id！", nil)
		return
	}

	filter := bson.M{
		"_id": objectId,
	}

	_, err = db.DeleteOne(collection, filter)
	if err != nil {
		resp.Error("删除失败！", nil)
		return
	}
	resp.Success("删除成功！", nil)
	return
}
