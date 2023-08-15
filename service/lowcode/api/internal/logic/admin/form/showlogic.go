package form

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

func (l *ShowLogic) Show(req *types.FormShowReq) (resp types.Response) {
	c := l.svcCtx

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(int64))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	form := model.Form{}
	var id primitive.ObjectID
	id, err = primitive.ObjectIDFromHex(req.FormId)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	filter := bson.M{"_id": id}
	err = form.Show(db, filter)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	db.Close()
	resp.Success("获取成功！", form)
	return
}
