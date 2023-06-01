package site

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.SiteShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId := req.SiteId
	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.ParseInt(userId.(string), 10, 64)
	var (
		db  database.MongoDB
		err error
	)
	db, err = c.MongoDB()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	// 查询当前用户是否包含当前站点
	site := new(model.Site)
	err = site.Show(db, bson.M{
		"siteId":            siteId,    // 匹配指定的文章ID
		"uid":               userIdInt, // 匹配指定的用户ID
		"isOwner":           1,
		"siteInfo.deleteAt": 0,
	})
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	filter := bson.M{"siteId": site.SiteId}

	// 设置更新内容
	unix := time.Now().Unix()
	update := bson.M{
		"$set": bson.M{
			"deleteAt": unix,
		},
	}

	site.DeleteAt = unix

	// 设置更新选项
	opts := options.Update().SetUpsert(false)
	collection := db.Collection("site")

	// 执行逻辑删除操作
	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		// 处理错误
		return
	}

	resp.Success("删除成功！", site)
	return
}
