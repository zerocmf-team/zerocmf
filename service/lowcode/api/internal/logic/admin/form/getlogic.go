package form

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.FormGetReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var forms []model.Form

	collection := db.Collection("form")

	// 查询条件
	filter := bson.M{} // 空的过滤条件表示查询所有记录
	sort := bson.D{{"listOrder", -1}}
	// 执行查询
	var cur *mongo.Cursor
	cur, err = collection.Find(context.TODO(), filter, options.Find().SetSort(sort))
	if err != nil {
		resp.Error("系统错误，查询失败", nil)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var form model.Form
		err = cur.Decode(&form)
		if err != nil {
			resp.Error("系统错误，查询失败", err.Error())
			return
		}
		forms = append(forms, form)
	}

	results := model.RecursionMenu(forms, primitive.ObjectID{0})

	resp.Success("获取成功！", results)
	return
}
