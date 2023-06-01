package site

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"zerocmf/common/bootstrap/database"
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

func (l *ShowLogic) Show(req *types.SiteShowReq) (resp types.Response) {
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
		"siteInfo.deleteAt": 0,
	})
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("操作成功！", site)
	return
}
