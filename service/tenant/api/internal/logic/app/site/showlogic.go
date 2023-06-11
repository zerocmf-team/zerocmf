package site

import (
	"context"
	"gorm.io/gorm"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/api/internal/svc"
	"zerocmf/service/tenant/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.SiteShowReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	site := model.Site{}

	tx := db.Where("site_id = ? and delete_at = 0", req.SiteId).First(&site)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			resp.Error("未查到该站点", nil)
			return
		}
		resp.Error("系统错误", tx.Error)
		return
	}
	resp.Success("获取成功！", site)
	return
}
