package formData

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"zerocmf/service/lowcode/model"

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

func (l *DeleteLogic) Delete(req *types.FormDataShowReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	formData := model.FormData{}
	var id primitive.ObjectID

	id, err = primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("formData")
	filter := bson.M{"_id": id}
	err = db.FindOne(collection, filter, &formData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error("该内容不存在", nil)
			return
		}
		resp.Error("find one err", err.Error())
		return
	}
	_, err = db.DeleteOne(collection, filter)
	if err != nil {
		resp.Error("删除失败", err.Error())
		return
	}
	resp.Success("删除成功！", formData)
	return
}
