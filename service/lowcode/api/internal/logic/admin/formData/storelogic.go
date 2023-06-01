package formData

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

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

func (l *StoreLogic) Store(req *types.FormDataSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.FormDataSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	form := model.Form{}
	var id primitive.ObjectID
	id, err = primitive.ObjectIDFromHex(req.FormId)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	filter := bson.M{"_id": id}
	err = form.Show(db, filter)
	if err != nil {
		resp.Error("form表单不存在", err.Error())
		return
	}

	/*	var schemaData model.Schema
		formSchema := form.Schema
		json.Unmarshal([]byte(formSchema), &schemaData)
		components := model.FindComponents(schemaData.ComponentsTree, "Form.Item")*/

	collection := db.Collection("formData")

	var bsonMap bson.M
	err = bson.UnmarshalExtJSON([]byte(req.FormDataJson), true, &bsonMap)
	if err != nil {
		resp.Error("ext json"+err.Error(), nil)
		return
	}

	if req.Id == "" {
		_, err = db.InsertOne(collection, bson.M{
			"schema": bsonMap,
		})
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		db.Close()
		resp.Success("操作成功！", bsonMap)
	} else {

	}
	return
}
