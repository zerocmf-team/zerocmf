package nav

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewEditLogic(header *http.Request, svcCtx *svc.ServiceContext) *EditLogic {
	ctx := header.Context()
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.NavSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	nav := model.Nav{}
	copier.Copy(&nav, &req)
	var tx *gorm.DB

	curNav := model.Nav{}
	tx = db.Where("id = ?", req.Id).First(&curNav)
	if util.IsDbErr(tx) != nil {
		resp.Error("系统错误", tx.Error)
		return
	}

	if curNav.Name != req.Name {
		showNav := model.Nav{}
		tx = db.Where("name = ?", req.Name).First(&showNav)
		if util.IsDbErr(tx) != nil {
			resp.Error("系统错误", tx.Error)
			return
		}
		if showNav.Id > 0 {
			resp.Error("该导航已存在！", tx.Error)
			return
		}
	}

	tx = db.Save(&nav)
	if tx.Error != nil {
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("更新成功！", nav)
	return
}
