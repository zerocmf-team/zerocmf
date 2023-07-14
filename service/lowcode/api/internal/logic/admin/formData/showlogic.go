package formData

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.FormDataShowReq) (resp types.Response) {

	c := l.svcCtx
	formId := req.Id
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	var objectID primitive.ObjectID
	objectID, err = primitive.ObjectIDFromHex(formId)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("formData")

	formData := model.FormData{}

	err = db.FindOne(collection, bson.M{"_id": objectID}, &formData)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	mapData := map[string]interface{}{}
	mapFormData := make(map[string]interface{}, 0)
	for _, kItem := range formData.Schema {
		fieldId := kItem.FieldId
		mapFormData[fieldId] = kItem.FieldData.Value
	}
	mapFormData["user"] = formData.User
	mapFormData["createTime"] = time.Unix(formData.CreateAt, 0).Format("2006-01-02 15:04:05")
	mapFormData["updateTime"] = time.Unix(formData.UpdateAt, 0).Format("2006-01-02 15:04:05")
	mapFormData["id"] = formData.Id
	mapData["formData"] = mapFormData
	mapData["formId"] = formData.FormId
	mapData["schema"] = formData.Schema

	resp.Success("获取成功！", mapData)
	return
}
