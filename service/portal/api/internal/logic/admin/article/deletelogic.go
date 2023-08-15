package article

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *DeleteLogic) Delete(req *types.ArticleDelReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id

	if id == 0 {
		resp.Error("分类id不能为空！", nil)
		return
	}

	portalPost := new(model.PortalPost)
	err := portalPost.Show(db, "id = ?", []interface{}{id})
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Error("该内容不存在", nil)
			return
		}

		resp.Error("操作失败", err.Error())
		return
	}
	portalPost.DeleteAt = time.Now().Unix()
	db.Where("id = ?", id).Updates(&portalPost)
	resp.Success("删除成功！", portalPost)
	return
}
