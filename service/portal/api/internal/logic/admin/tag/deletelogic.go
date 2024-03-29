package tag

import (
	"context"
	"errors"
	"net/http"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(header *http.Request, svcCtx *svc.ServiceContext) *DeleteLogic {
	ctx := header.Context()
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
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

	resp.Success("删除成功！", nil)
	return

}
