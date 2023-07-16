package settings

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

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

func (l *StoreLogic) Store(req *types.SettingSaveReq) (resp types.Response) {
	resp = save(l.svcCtx, req)
	return
}

func save(c *svc.ServiceContext, req *types.SettingSaveReq) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	settings := model.Settings{
		Key: req.Key,
	}

	var saveData = bson.M{}

	params := make(map[string]interface{})

	json.Unmarshal([]byte(req.FormDataJson), &params)

	saveData, err = settings.Store(db, params)

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("操作成功！", saveData)
	return
}
