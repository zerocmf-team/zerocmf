/**
Desc: 显示全部菜单
Author: daifuyang
Contact: github.com/daifuyang
Date: Date: 2023-07-06 19:49:24
*/

package adminMenu

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (l *GetLogic) Get(req *types.AdminMenuGetReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	plugin := req.Plugin
	filter := bson.M{}

	if plugin != "" {
		filter["plugin"] = plugin
	}

	adminMenu := model.AdminMenu{}
	menus, mErr := adminMenu.GetTrees(db, l.ctx, filter)
	if mErr != nil && !errors.Is(mErr, mongo.ErrNoDocuments) {
		resp.Error("菜单获取失败！", mErr.Error())
		return
	}

	resp.Success("获取成功！", menus)
	return
}
