package account

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

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

func (l *GetLogic) Get(req *types.UserAdminListReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	// 计算要跳过的文档数量
	skip := (current - 1) * pageSize

	match := bson.M{
		"deletedAt": 0,
	}

	if req.UserLogin != "" {
		regex := primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*$", req.UserLogin), Options: "i"}
		match["userLogin"] = bson.M{"$regex": regex}
	}

	if req.UserEmail != "" {
		regex := primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*$", req.UserEmail), Options: "i"}
		match["userEmail"] = bson.M{"$regex": regex}
	}

	// 创建聚合管道
	pipeline := []bson.M{
		{
			"$match": match,
		},
		{
			"$facet": bson.M{
				"data": []bson.M{
					{
						"$skip": skip,
					},
					{
						"$limit": pageSize,
					},
				},
				"totalCount": []bson.M{
					{
						"$count": "total",
					},
				},
			},
		},
	}

	collection := db.Collection("user")

	// 执行查询
	var cur *mongo.Cursor
	cur, err = collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		resp.Error("系统错误，查询失败", nil)
		return
	}
	defer cur.Close(context.TODO())

	paginate := data.Paginate{}

	if cur.Next(context.TODO()) {
		var aggregate struct {
			Data       []model.User `bson:"data"`
			TotalCount []struct {
				Total int `bson:"total"`
			} `bson:"totalCount"`
		}
		err = cur.Decode(&aggregate)
		if err != nil {
			resp.Error("系统错误，查询失败", err.Error())
			return
		}

		var userData = make([]model.User, 0)
		for _, v := range aggregate.Data {
			if v.LastLoginAt > 0 {
				v.LastLoginTime = time.Unix(v.LastLoginAt, 0).Format("2006-01-02 15:04:05")
			}
			if v.CreateAt > 0 {
				v.CreateTime = time.Unix(v.CreateAt, 0).Format("2006-01-02 15:04:05")
			}
			if v.UpdateAt > 0 {
				v.UpdateTime = time.Unix(v.UpdateAt, 0).Format("2006-01-02 15:04:05")
			}
			userData = append(userData, v)
		}

		total := 0
		totalCount := aggregate.TotalCount
		if len(totalCount) > 0 {
			total = totalCount[0].Total
		}

		paginate = data.Paginate{
			Data:     userData,
			Current:  current,
			PageSize: pageSize,
			Total:    int64(total),
		}

	}

	resp.Success("获取成功！", paginate)
	return
}
