package department

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

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

func (l *EditLogic) Edit(req *types.DepReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	id := req.Id
	if id <= 0 {
		resp.Error("参数不合法", nil)
		return
	}

	department := model.Department{}
	tx := db.Where("id = ?", id).First(&department)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			resp.Error("该部门不存在！", nil)
			return
		}
		resp.Error("系统错误", nil)
		return
	}

	return saveDepartment(req, c)
}
