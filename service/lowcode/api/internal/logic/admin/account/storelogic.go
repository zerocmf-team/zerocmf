package account

import (
	"context"
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

func (l *StoreLogic) Store(req *types.UserAdminSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.UserAdminSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	collection := db.Collection("user")
	userShow := model.User{}
	err = db.FindOne(collection, bson.M{
		"userLogin": req.UserLogin,
		"deletedAt": time.Now().Unix(),
	}, &userShow)

	if err != nil && err != mongo.ErrNoDocuments {
		resp.Error(err.Error(), nil)
		return
	}

	if req.UserLogin == "" {
		resp.Error("用户名不能为空！", nil)
		return
	}

	user := model.User{
		UserType: 1,
	}
	if req.Id == "" {
		if !userShow.Id.IsZero() {
			resp.Error("该账号名已存在！", nil)
			return
		}
		copier.Copy(&user, &req)

		user.CreateAt = time.Now().Unix()

		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, &user)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		user.Id = one.InsertedID.(primitive.ObjectID)

	} else {

		var objectID primitive.ObjectID
		objectID, err = primitive.ObjectIDFromHex(req.Id)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		filter := bson.M{
			"_id": objectID,
		}

		err = db.FindOne(collection, &filter, &user)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		// 不是同一条字段
		if !userShow.Id.IsZero() && user.Id != userShow.Id {
			resp.Error("该角色名称已存在", nil)
			return
		}

		copier.Copy(&user, &req)

		user.UpdateAt = time.Now().Unix()

		update := bson.M{
			"$set": &user,
		}

		_, err = db.UpdateOne(collection, filter, update)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

	}

	resp.Success("操作成功！", user)
	return
}
