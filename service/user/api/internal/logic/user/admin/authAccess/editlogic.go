package authAccess

import (
	"context"
	"github.com/jinzhu/copier"

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

func (l *EditLogic) Edit(req *types.AccessEdit) (resp types.Response) {
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
