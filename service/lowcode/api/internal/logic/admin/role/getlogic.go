package role

import (
	"context"
	"fmt"
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

func (l *GetLogic) Get(req *types.UserAdminRoleList) (resp types.Response) {
	c := l.svcCtx
	r := c.Request
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var hasPagination = make([]bson.M, 0)

	noPagination := 0
	if req.NoPagination != nil {
		noPagination = *req.NoPagination
	}

	paginate := data.Paginate{}

	if noPagination != 1 {
		var (
			current  int
			pageSize int
		)
		current, pageSize, err = data.NewPaginate(r).Default()
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

		paginate.Current = current
		paginate.PageSize = pageSize

		// 计算要跳过的文档数量
		skip := (current - 1) * pageSize

		hasPagination = []bson.M{
			{
				"$skip": skip,
			},
			{
				"$limit": pageSize,
			},
		}
	}

	// 创建聚合管道
	match := bson.M{
		"deletedAt": 0,
	}

	if req.Name != "" {
		regex := primitive.Regex{Pattern: fmt.Sprintf("^.*%s.*$", req.Name), Options: "i"}
		match["name"] = bson.M{"$regex": regex}
	}

	if req.Status != nil {
		match["status"] = req.Status
	}

	pipeline := []bson.M{
		{
			"$match": match,
		},
		{
			"$facet": bson.M{
				"data": hasPagination,
				"totalCount": []bson.M{
					{
						"$count": "total",
					},
				},
			},
		},
	}

	collection := db.Collection("role")

	// 执行查询
	var cur *mongo.Cursor
	cur, err = collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		resp.Error("系统错误，查询失败", nil)
		return
	}
	defer cur.Close(context.TODO())

	var roleData = make([]model.Role, 0)
	if cur.Next(context.TODO()) {
		var aggregate struct {
			Data       []model.Role `bson:"data"`
			TotalCount []struct {
				Total int `bson:"total"`
			} `bson:"totalCount"`
		}
		err = cur.Decode(&aggregate)
		if err != nil {
			resp.Error("系统错误，查询失败", err.Error())
			return
		}

		for _, v := range aggregate.Data {
			if v.CreateAt > 0 {
				v.CreateTime = time.Unix(v.CreateAt, 0).Format("2006-01-02 15:04:05")
			}
			if v.UpdateAt > 0 {
				v.UpdateTime = time.Unix(v.UpdateAt, 0).Format("2006-01-02 15:04:05")
			}
			roleData = append(roleData, v)
		}

		total := 0
		totalCount := aggregate.TotalCount
		if len(totalCount) > 0 {
			total = totalCount[0].Total
		}

		paginate.Data = roleData
		paginate.Total = int64(total)
	}

	if noPagination == 1 {
		resp.Success("获取成功！", roleData)
		return
	}

	resp.Success("获取成功！", paginate)
	return
}
