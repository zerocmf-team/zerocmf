package form

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"time"
	"zerocmf/service/portal/model"

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

func (l *StoreLogic) Store(req *types.FormSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	resp = saveForm(db, req, 0)
	return
}

func saveForm(db *gorm.DB, req *types.FormSaveReq, typ int) (resp types.Response) {
	form := new(model.Form)
	copier.Copy(&form, &req)
	var tx *gorm.DB
	if typ == 0 {
		form.CreateAt = time.Now().Unix()
		tx = db.Create(&form)
	} else {
		form.UpdateAt = time.Now().Unix()
		tx = db.Save(&form)
	}

	if tx.Error != nil {
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("操作成功", form)
	return
}
