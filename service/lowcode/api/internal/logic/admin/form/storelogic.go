package form

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
	bsModel "zerocmf/common/bootstrap/model"
	"zerocmf/common/bootstrap/util"
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

func (l *StoreLogic) Store(req *types.FormSaveReq) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.FormSaveReq) (resp types.Response) {

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.ParseInt(userId.(string), 10, 64)
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var schemaData model.Schema
	formSchema := req.Schema
	json.Unmarshal([]byte(formSchema), &schemaData)
	components := model.FindComponents(schemaData.ComponentsTree, "Form.Item")

	var rules []model.SRules
	for _, component := range components {
		_props := component.Props
		propsBytes, _ := json.Marshal(_props)
		var props model.SProps
		json.Unmarshal(propsBytes, &props)
		if len(props.Rules) > 0 {
			rules = append(rules, props.Rules...)
		}
	}

	collection := db.Collection("form")

	if req.Id == "" {
		//新增表单
		form := model.Form{
			ParentId:  primitive.ObjectID{},
			UserId:    userIdInt,
			ListOrder: 10000,
			Status:    1,
			Time: bsModel.Time{
				CreateAt: time.Now().Unix(),
			},
		}
		copier.Copy(&form, &req)
		var parentId primitive.ObjectID
		parentId, err = primitive.ObjectIDFromHex(req.ParentId)
		form.ParentId = parentId
		//	选择表单
		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, form)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		form.Id = one.InsertedID.(primitive.ObjectID)
		db.Close()
		resp.Success("操作成功！", form)
	} else {
		form := model.Form{}
		var id primitive.ObjectID
		id, err = primitive.ObjectIDFromHex(req.Id)
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
		form.Rules = rules
		form.UpdateAt = time.Now().Unix()
		copier.Copy(&form, &req)
		var bsonM bson.M
		bsonM, err = util.AtoBsonM(form)
		if err != nil {
			resp.Error("AtoBsonM err", err.Error())
			return
		}

		// 设置更新内容
		update := bson.M{
			"$set": bsonM,
		}

		// 设置更新选项
		opts := options.Update().SetUpsert(false)
		collection = db.Collection("form")

		// 执行逻辑删除操作
		_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			// 处理错误
			return
		}

		resp.Success("更新成功", form)

	}
	return
}
