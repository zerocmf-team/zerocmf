package app_page

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

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

func (l *StoreLogic) Store(req *types.AppPageSaveReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	resp = savePage(db, req, 0)
	return
}

func savePage(db *gorm.DB, req *types.AppPageSaveReq, typ int) (resp types.Response) {
	appPage := new(model.AppPage)
	appPage.Status = 1
	copier.Copy(&appPage, &req)
	var tx *gorm.DB
	if typ == 0 {
		appPage.CreateAt = time.Now().Unix()
		tx = db.Create(&appPage)
	} else {
		appPage.UpdateAt = time.Now().Unix()
		tx = db.Save(&appPage)
	}

	if tx.Error != nil {
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("操作成功", appPage)
	return
}
