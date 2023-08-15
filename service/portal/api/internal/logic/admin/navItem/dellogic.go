package navItem

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDelLogic(header *http.Request, svcCtx *svc.ServiceContext) *DelLogic {
	ctx := header.Context()
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	id := req.Id

	// 查询是否存在子分类
	var navItem []model.NavItem
	tx := db.Where("parent_id = ?", id).Find(&navItem)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if tx.RowsAffected > 0 {
		resp.Error("请先删除子分类！", nil)
		return
	}

	tx = db.Where("id = ?", id).Delete(&navItem)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success("删除成功！", nil)
	return
}
