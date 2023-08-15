package themeFile

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(header *http.Request, svcCtx *svc.ServiceContext) *SaveLogic {
	ctx := header.Context()
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req *types.ThemeFileSaveReq) (resp types.Response) {

	c := l.svcCtx
	id := req.Id
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	more := req.More

	if more == "" {
		resp.Error("配置不能为空！", nil)
		return
	}

	themeFile := model.ThemeFile{}
	tx := db.Where("id = ?", id).First(&themeFile)
	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	tx = db.Model(&model.ThemeFile{}).Where("id = ?", id).Update("more", more)
	if tx.Error != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}

	resp.Success("更新成功！", more)
	return
}
