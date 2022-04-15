package tag

import (
	"context"
	"errors"
	"gincmf/common/bootstrap/util"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

func (l *DeleteLogic) Delete(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	id := req.Id

	var tag model.PortalTag
	tx := db.Where("id = ?", id).First(&tag)

	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		resp.Error("该文章不存在或已删除", nil)
		return
	}

	if tx.RowsAffected == 0 {
		resp.Error("内容不存在！", nil)
		return
	}

	tx = db.Where("id", id).Delete(&tag)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success( "删除成功！", nil)
	return

}
