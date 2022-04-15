package authAccess

import (
	"context"
	"github.com/jinzhu/copier"

	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"

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

func (l *EditLogic) Edit(req *types.AccessEdit) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	form := access{}
	copier.Copy(&form, &req)
	role, err := save(form, l.svcCtx)

	if err != nil {
		resp.Error(err.Error(),nil)
		return
	}
	resp.Success("操作成功！",role)
	return
}
