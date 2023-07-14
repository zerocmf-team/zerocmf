package settings

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.SettingSaveReq) (resp types.Response) {
	resp = save(l.svcCtx, req)
	return
}

func save(c *svc.ServiceContext, req *types.SettingSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("settings")
	// 新增

	filter := bson.M{"key": req.Key}
	result := bson.M{}
	err = db.FindOne(collection, filter, &result)

	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			resp.Error("查询失败", nil)
			return
		}
	}

	var objectId primitive.ObjectID
	id := result["_id"]
	if id != nil {
		objectId = id.(primitive.ObjectID)
	}

	var params map[string]interface{}
	err = json.Unmarshal([]byte(req.FormDataJson), &params)
	if err != nil {
		resp.Error("参数不合法", err.Error())
		return
	}

	saveData := bson.M{
		"key":   req.Key,
		"value": params,
	}

	if objectId.IsZero() {
		//var one *mongo.InsertOneResult
		_, err = db.InsertOne(collection, &saveData)
		if err != nil {
			resp.Error("新增失败", err.Error())
			return
		}
	} else {
		_, err = db.UpdateOne(collection, filter, bson.M{
			"$set": saveData,
		})
		if err != nil {
			resp.Error("更新失败", err.Error())
			return
		}
	}
	resp.Success("操作成功！", saveData)
	return
}
