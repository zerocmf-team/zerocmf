package themePage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (l *ShowLogic) Show(req *types.ThemePageShowReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	id := req.Id

	if id == "" {
		resp.Error("id不能为空！", nil)
		return
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	//if objectErr != nil {
	//	resp.Error("非法id！", nil)
	//	return
	//}

	filter := bson.M{
		"_id": objectId,
	}
	page := model.ThemePage{}

	if !objectId.IsZero() {
		err = page.Show(db, filter)
		if err != nil {
			resp.Error("查询错误！", err.Error())
			return
		}

	}

	if page.Id.IsZero() {
		typ := req.Type
		themeKey := req.ThemeKey

		if themeKey == "" {
			resp.Error("主题不能为空！", err.Error())
			return
		}

		filter = bson.M{
			"themeKey": themeKey,
			"key":      typ,
		}
		err = page.Show(db, filter)
		if err != nil {
			resp.Error("查询错误！", err.Error())
			return
		}
	}

	resp.Success("获取成功！", page)

	return
}
