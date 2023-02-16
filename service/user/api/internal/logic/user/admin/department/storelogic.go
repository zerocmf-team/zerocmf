package department

import (
	"context"
	"github.com/jinzhu/copier"
	"time"
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

func (l *StoreLogic) Store(req *types.DepReq) (resp types.Response) {
	c := l.svcCtx
	return saveDepartment(req, c)
}

func saveDepartment(req *types.DepReq, c *svc.ServiceContext) (resp types.Response) {
	dep := model.Department{}
	copier.Copy(&dep, &req)

	if req.Id == 0 {
		dep.CreateAt = time.Now().Unix()
	}else {
		dep.UpdateAt = time.Now().Unix()
	}

	tx := c.Db.Save(&dep)
	if tx.Error != nil {
		resp.Error("系统异常", tx.Error.Error())
		return
	}
	resp.Success("操作成功", dep)
	return
}
