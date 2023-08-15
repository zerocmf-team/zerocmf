package app_page

import (
	"context"
	"net/http"
	"time"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

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

func (l *StoreLogic) Store(req *types.AppPageSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	resp = savePage(db, req, 0)
	return
}

func savePage(db *gorm.DB, req *types.AppPageSaveReq, typ int) (resp types.Response) {
	appPage := new(model.AppPage)
	copier.Copy(&appPage, &req)
	var tx *gorm.DB

	// 查询公共头是否存在
	if req.Type == "1" {
		page := new(model.AppPage)
		tx = db.Where("is_public = ? AND name = ?", req.IsPublic, req.Name).First(&page)
		if tx.Error != nil {
			resp.Error("系统错误", tx.Error)
			return
		}
		if page.Id != 0 {
			appPage.Id = page.Id
			typ = 1
		}
	}

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
