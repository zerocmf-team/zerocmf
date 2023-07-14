package themePage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/service/lowcode/model"

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

func (l *StoreLogic) Store(req *types.ThemePageSaveReq) (resp types.Response) {
	resp = save(l.svcCtx, req)
	return
}

func save(c *svc.ServiceContext, req *types.ThemePageSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	filter := bson.M{
		"name": req.Name,
	}

	collection := db.Collection("themePage")

	themePage := model.ThemePage{}
	findErr := db.FindOne(collection, filter, &themePage)
	if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
		resp.Error(findErr.Error(), nil)
		return
	}

	// 新增
	if req.Id == "" {

		if !themePage.Id.IsZero() {
			resp.Error("该页面已经存在！", nil)
			return
		}

		themePage = model.ThemePage{
			ThemeKey:    req.ThemeKey,
			Name:        req.Name,
			Description: req.Description,
			Type:        req.Type,
			IsPublic:    req.IsPublic,
			Schema:      req.Schema,
			ListOrder:   req.ListOrder,
			UserId:      1,
			CreateAt:    time.Now().Unix(),
			UpdateAt:    time.Now().Unix(),
			Status:      1,
			DeleteAt:    0,
		}

		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, &themePage)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		themePage.Id = one.InsertedID.(primitive.ObjectID)

		resp.Success("新增成功！", themePage)

	} else {

		themePage.ThemeKey = req.ThemeKey
		themePage.Name = req.Name
		themePage.Description = req.Description
		themePage.IsPublic = req.IsPublic
		themePage.Type = req.Type
		themePage.Schema = req.Schema
		themePage.ListOrder = req.ListOrder
		themePage.UserId = 1
		themePage.CreateAt = time.Now().Unix()
		themePage.UpdateAt = time.Now().Unix()
		themePage.Status = 1
		themePage.DeleteAt = 0

		_, err = db.UpdateOne(collection, filter, bson.M{
			"$set": themePage,
		})
		if err != nil {
			resp.Error("更新失败", err.Error())
			return
		}
		resp.Success("获取成功！", themePage)
	}

	return
}
