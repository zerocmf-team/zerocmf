package app

import (
	"context"
	"net/http"
	"time"
	"zerocmf/service/portal/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(header *http.Request, svcCtx *svc.ServiceContext) *StoreLogic {
	ctx := header.Context()
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.AppSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	resp = saveApp(db, req, 0)
	return
}

func saveApp(db *gorm.DB, req *types.AppSaveReq, typ int) (resp types.Response) {
	var tx *gorm.DB
	app := new(model.App)
	copier.Copy(&app, &req)
	if typ == 0 {
		now := time.Now().Unix()
		app.CreateAt = now
		tx = db.Create(&app)
		// 创建首页页面
		if tx.Error != nil {
			appPage := model.AppPage{
				AppId:       app.Id,
				IsHome:      1,
				Name:        "首页",
				Description: "首页",
				Type:        "page",
				CreateAt:    now,
				UpdateAt:    now,
			}
			tx = db.Create(&appPage)
		}

	} else {
		app.UpdateAt = time.Now().Unix()
		tx = db.Save(&app)
	}
	if tx.Error != nil {
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("操作成功", nil)
	return
}
