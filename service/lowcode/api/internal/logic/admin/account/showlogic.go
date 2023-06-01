package account

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"
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

func (l *ShowLogic) Show(req *types.UserAdminShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	collection := db.Collection("user")

	var objectID primitive.ObjectID
	objectID, err = primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	filter := bson.M{
		"_id":       objectID,
		"deletedAt": 0,
	}

	user := model.User{}

	err = db.FindOne(collection, filter, &user)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", user)
	return
}
