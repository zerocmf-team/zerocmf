package nav

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.NavSaveReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
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
