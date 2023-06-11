package nav

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"zerocmf/common/bootstrap/util"
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

func (l *StoreLogic) Store(req *types.NavSaveReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	nav := model.Nav{}
	copier.Copy(&nav, &req)
	var tx *gorm.DB
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
	tx = db.Create(&nav)
	if tx.Error != nil {
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("新增成功！", nav)
	return
}
