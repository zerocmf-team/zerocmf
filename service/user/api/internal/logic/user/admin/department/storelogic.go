package department

import (
	"context"
	"github.com/jinzhu/copier"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.DepReq) (resp *types.Response) {
	c := l.svcCtx
	return saveDepartment(req, c)
}

func saveDepartment(req *types.DepReq, c *svc.ServiceContext) (resp *types.Response) {
	resp = new(types.Response)
	dep := model.Department{}
	copier.Copy(&dep, &req)
	tx := c.Db.Save(&dep)
	if tx.Error != nil {
		resp.Error("系统异常", tx.Error.Error())
		return
	}
	resp.Success("操作成功", dep)
	return
}
