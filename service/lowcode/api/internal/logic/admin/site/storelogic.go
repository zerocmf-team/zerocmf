package site

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/database"
	bsModel "zerocmf/common/bootstrap/model"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"
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

func (l *StoreLogic) Store(req *types.SiteSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.SiteSaveReq) (resp types.Response) {

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.ParseInt(userId.(string), 10, 64)

	db, err := c.MongoDB()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	redis := c.Redis()

	site := model.Site{
		SiteId:   0,
		Name:     "",
		Desc:     "",
		Status:   1,
		DeleteAt: 0,
		Time: bsModel.Time{
			CreateAt: time.Now().Unix(),
		},
	}
	copier.Copy(&site, &req)
	collection := db.Collection("site")
	if req.SiteId == 0 {
		// 插入一条
		var siteId int64
		key := "lowcode:site"
		siteId, err = redis.EncryptUid(key, 0)
		if err != nil {
			resp.Error("操作失败！", err.Error())
			return
		}
		site.SiteId = siteId

		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, site)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		site.Id = one.InsertedID.(primitive.ObjectID)

		// 插入一条管理员信息

		// 选择租户表
		var tenantDb database.MongoDB
		tenantDb, err = c.MongoDB("tenant_" + strconv.FormatInt(siteId, 10))
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		user := model.User{
			UserType:    1,
			Root:        1,
			UserLogin:   "admin",
			UserPass:    util.GetMd5("123456"),
			LastLoginAt: time.Now().Unix(),
			Time: bsModel.Time{
				CreateAt: time.Now().Unix(),
			},
		}

		userCollection := tenantDb.Collection("user")
		one, err = tenantDb.InsertOne(userCollection, user)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		new(model.Role).AutoMigrate(tenantDb)
		tenantDb.Close()

		// 建立关系表
		suCollection := db.Collection("siteUser")
		siteUser := model.SiteUser{
			SiteId:    siteId,
			Uid:       userIdInt,
			Oid:       one.InsertedID.(primitive.ObjectID),
			IsOwner:   1,
			ListOrder: 10000,
			Status:    1,
		}
		one, err = db.InsertOne(suCollection, siteUser)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		db.Close()

		resp.Success("操作成功！", site)
	} else {
		site = model.Site{}
		err = site.Show(db, bson.M{
			"siteId":            req.SiteId,
			"uid":               userIdInt,
			"isOwner":           1,
			"siteInfo.deleteAt": 0,
		})
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		copier.Copy(&site, &req)
		site.UpdateAt = time.Now().Unix()

		var bsonM bson.M
		bsonM, err = util.AtoBsonM(site)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		filter := bson.M{"_id": site.Id}

		// 设置更新内容
		update := bson.M{
			"$set": bsonM,
		}

		// 设置更新选项
		opts := options.Update().SetUpsert(false)
		collection = db.Collection("site")

		// 执行逻辑删除操作
		_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			// 处理错误
			return
		}

		resp.Success("更新成功", site)

	}
	return
}
