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

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DepOneReq) (resp *types.Response) {
	resp = new(types.Response)
	c := l.svcCtx
	db := c.Db
	department := model.Department{}
	id := req.Id
	if  id <= 0 {
		resp.Error("参数不合法",nil)
		return
	}
	tx :=db.Where("id = ?",id).First(&department)
	if tx.Error != nil {
		if errors.Is(tx.Error,gorm.ErrRecordNotFound) {
			resp.Error("该部门不存在！",nil)
			return
		}
		resp.Error("系统错误",nil)
		return
	}

	var count int64 = 0
	tx = db.Where("parent_id",id).Model(&department).Count(&count)
	if tx.Error != nil {
		resp.Error("系统错误",nil)
		return
	}
	if count > 0 {
		resp.Error("请先删除子集菜单！",nil)
		return
	}
	
	tx = db.Where("id = ?",id).Delete(&department)
	if tx.Error != nil {
		resp.Error("系统错误",nil)
		return
	}
	resp.Success("删除成功！",department)
	return
}
