package site

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
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

func (l *GetLogic) Get(req *types.SiteGetReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.ParseInt(userId.(string), 10, 64)

	db, err := c.MongoDB()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("siteUser")

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	skip := (current - 1) * pageSize

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "site",
				"localField":   "siteId",
				"foreignField": "siteId",
				"as":           "siteInfo",
			},
		},
		{
			"$match": bson.M{
				"uid":               userIdInt,
				"siteInfo.deleteAt": 0,
			},
		},
		{
			"$unwind": "$siteInfo", // 平铺数组字段
		},
		{
			"$project": bson.M{

				"oid":       1,
				"isOwner":   1,
				"createAt":  1,
				"listOrder": 1,
				"status":    1,
				"siteId":    "$siteInfo.siteId",
				"name":      "$siteInfo.name",
				"desc":      "$siteInfo.desc",
				"domain":    "$siteInfo.domain",
			},
		},
		{
			"$facet": bson.M{
				"data": []bson.M{
					{"$skip": skip},
					{"$limit": pageSize},
				},
				"total": []bson.M{
					{"$count": "count"},
				},
			},
		},
	}
	var cursor *mongo.Cursor
	cursor, err = collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		resp.Error("操作失败！", err.Error())
		return
	}
	defer cursor.Close(context.Background())

	paginate := data.Paginate{}

	// 检查结果集是否为空

	var total int64

	if cursor.Next(context.Background()) {
		// 解析结果
		type Result struct {
			Data  []bson.M `bson:"data"`
			Total []struct {
				Count int `bson:"count"`
			} `bson:"total"`
		}

		var result Result
		if err = cursor.Decode(&result); err != nil {
			resp.Error("操作失败！", err.Error())
			return
		}

		if len(result.Total) > 0 {
			totalRecords := result.Total[0].Count
			total = int64(totalRecords)
		}

		paginate = data.Paginate{
			Data:     result.Data,
			Current:  current,
			PageSize: pageSize,
			Total:    total,
		}
	}

	db.Close()
	resp.Success("操作成功！", paginate)
	return
}
