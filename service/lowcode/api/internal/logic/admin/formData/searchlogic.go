package formData

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.FormSearchReq) (resp types.Response) {
	c := l.svcCtx
	formId := req.FormId
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

	current := 1
	pageSize := 10

	if req.Current != nil {
		current = *req.Current
	}

	if req.PageSize != nil {
		pageSize = *req.PageSize
	}

	filter := bson.M{
		"formId": objectID,
	} // 空的过滤条件表示查询所有记录

	var searchFieldJson = make(map[string]interface{}, 0)
	if req.SearchFieldJson != "" {
		err = json.Unmarshal([]byte(req.SearchFieldJson), &searchFieldJson)
		if err != nil {
			resp.Error("json Unmarshal "+err.Error(), nil)
			return
		}
	}

	for k, v := range searchFieldJson {
		filter["schema.fieldId"] = k
		filter["schema.fieldData.value"] = v
	}

	pipeline := []bson.M{
		{
			"$match": filter,
		},
	}

	if pageSize > 0 {
		skip := (current - 1) * pageSize
		pipeline = append(pipeline, bson.M{
			"$facet": bson.M{
				"data": []bson.M{
					{
						"$sort": bson.M{
							"createAt": -1,
						},
					},
					{
						"$skip": skip,
					},
					{
						"$limit": pageSize,
					},
				},
				"pagination": []bson.M{
					{"$count": "total"},
				},
			},
		})
	} else {
		pipeline = append(pipeline, bson.M{
			"$facet": bson.M{
				"data": []bson.M{
					{
						"$sort": bson.M{
							"createAt": -1,
						},
					},
				},
			},
		})
	}

	// 执行查询
	var cur *mongo.Cursor
	cur, err = collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		resp.Error("系统错误，查询失败", nil)
		return
	}

	defer cur.Close(context.TODO())

	formData := make([]map[string]interface{}, 0)

	var result struct {
		Data       []model.FormData `bson:"data" json:"data"`
		Pagination []struct {
			Total int64 `bson:"total" json:"total"`
		} `bson:"pagination" json:"pagination"`
	}

	paginate := data.Paginate{
		Current:  current,
		PageSize: pageSize,
	}

	if cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			resp.Error("解析错误", err.Error())
			return
		}

		data := result.Data

		for _, item := range data {
			mapData := map[string]interface{}{}
			mapFormData := make(map[string]interface{}, 0)
			for _, kItem := range item.Schema {
				fieldId := kItem.FieldId
				mapFormData[fieldId] = kItem.FieldData.Value
				mapFormData["user"] = item.User
				mapFormData["createTime"] = time.Unix(item.CreateAt, 0).Format("2006-01-02 15:04:05")
				mapFormData["updateTime"] = time.Unix(item.UpdateAt, 0).Format("2006-01-02 15:04:05")
			}
			mapFormData["id"] = item.Id
			mapData["formData"] = mapFormData
			mapData["formId"] = item.FormId
			mapData["schema"] = item.Schema
			formData = append(formData, mapData)
		}
		paginate.Data = formData

		if len(result.Pagination) > 0 {
			total := result.Pagination[0].Total
			paginate.Total = total
		}

	}

	if pageSize == 0 {
		resp.Success("获取成功！", formData)
		return
	}

	resp.Success("获取成功！", paginate)
	return
}
