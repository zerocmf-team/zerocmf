package themePage

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
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
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("themePage")

	// 新增
	if req.Id == "" {
		filter := bson.M{
			"name": req.Name,
		}
		themePage := model.ThemePage{}
		findErr := db.FindOne(collection, filter, &themePage)
		if findErr != nil && !errors.Is(findErr, mongo.ErrNoDocuments) {
			resp.Error(findErr.Error(), nil)
			return
		}

		if !themePage.Id.IsZero() {
			resp.Error("该页面已经存在！", nil)
			return
		}

		themePage = model.ThemePage{
			Theme:       req.Theme,
			ThemeType:   req.ThemeType,
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

		var objectId primitive.ObjectID
		objectId, err = primitive.ObjectIDFromHex(req.Id)
		if err != nil {
			resp.Error("非法请求", nil)
			return
		}

		filter := bson.M{
			"_id": objectId,
		}

		themePage := model.ThemePage{}
		findErr := db.FindOne(collection, filter, &themePage)
		if findErr != nil {
			resp.Error(findErr.Error(), nil)
			return
		}

		copier.CopyWithOption(&themePage, &req, copier.Option{IgnoreEmpty: true})
		//themePage.Theme = req.Theme
		//themePage.ThemeType = req.ThemeType
		//themePage.Name = req.Name
		//themePage.Description = req.Description
		//themePage.IsPublic = req.IsPublic
		//themePage.Type = req.Type
		//themePage.Schema = req.Schema
		//themePage.ListOrder = req.ListOrder
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
		resp.Success("更新成功！", themePage)
	}

	return
}
