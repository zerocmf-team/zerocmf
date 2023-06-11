package formData

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"
	"zerocmf/service/user/rpc/userclient"

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

	h := c.Request
	userId, _ := c.Get("userId")
	siteId, _ := c.Get("siteId")

	reply, err := c.UserRpc.Get(h.Context(), &userclient.UserRequest{
		UserId: userId.(string),
		SiteId: siteId.(string),
	})
	if err != nil {
		resp.Error("获取站点用户失败！", err.Error())
		return
	}

	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	form := model.Form{}
	var formId primitive.ObjectID
	formId, err = primitive.ObjectIDFromHex(req.FormId)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	filter := bson.M{"_id": formId}
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

	var schema map[string]interface{}

	err = json.Unmarshal([]byte(req.FormDataJson), &schema)
	if err != nil {
		resp.Error("参数不合法", err.Error())
		return
	}

	columns := form.Columns

	for ck, cv := range columns {
		// todo rules 规则校验
		inVal := schema[cv.FieldId]
		if columns[ck].FieldData == nil {
			columns[ck].FieldData = new(model.FieldData)
		}
		columns[ck].FieldData.Text = cv.Label
		columns[ck].FieldData.Value = inVal
	}

	//err = bson.UnmarshalExtJSON([]byte(req.FormDataJson), true, &schema)
	//if err != nil {
	//	resp.Error("ext json"+err.Error(), nil)
	//	return
	//}

	formData := model.FormData{
		FormId:    formId,
		Schema:    columns,
		UserId:    reply.Id,
		UserLogin: reply.UserLogin,
	}

	if req.Id == "" {
		formData.CreateAt = time.Now().Unix()
		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, formData)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		formData.Id = one.InsertedID.(primitive.ObjectID)

	} else {

		formData.UpdateAt = time.Now().Unix()

		// 设置更新内容
		update := bson.M{
			"$set": formData,
		}

		var id primitive.ObjectID
		id, err = primitive.ObjectIDFromHex(req.Id)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		formData.Id = id

		filter = bson.M{"_id": id}

		var _ *mongo.UpdateResult
		_, err = db.UpdateOne(collection, filter, update)
		if err != nil {
			resp.Error("更新失败", err.Error())
			return
		}

	}
	db.Close()
	resp.Success("操作成功！", formData)
	return
}
