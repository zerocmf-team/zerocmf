package settings

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.SettingShowReq) (resp types.Response) {

	c := l.svcCtx

	siteId, _ := c.Get("siteId")

	key := req.Key

	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("settings")
	// 新增

	filter := bson.M{"key": key}

	var result = bson.M{}

	err = db.FindOne(collection, &filter, &result)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		resp.Error("查询失败", nil)
		return
	}

	if result["value"] == nil {
		result["value"] = bson.M{}
	}

	resp.Success("获取成功！", result["value"])
	return
}
