package account

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/service/user/model"

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

func (l *DeleteLogic) Delete(req *types.UserAdminShowReq) (resp types.Response) {

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
		"_id": objectID,
	}

	user := model.User{}

	err = db.FindOne(collection, &filter, &user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			resp.Error("该管理员不存在！", nil)
			return
		}
		resp.Error(err.Error(), nil)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"deletedAt": time.Now().Unix(),
		},
	}

	_, err = db.UpdateOne(collection, filter, update)
	if err != nil {
		resp.Error(err.Error(), nil)
	}

	resp.Success("删除成功！", user)

	return
}
