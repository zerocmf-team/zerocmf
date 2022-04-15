package themeFile

import (
	"context"
	"gincmf/common/bootstrap/util"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req *types.ThemeFileSaveReq) (resp types.Response) {

	c := l.svcCtx
	id := req.Id
	db := c.Db

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
