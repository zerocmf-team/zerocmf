package role

import (
	"context"
	"strconv"
	"zerocmf/common/bootstrap/casbin"
	"zerocmf/common/bootstrap/util"
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

func (l *DeleteLogic) Delete(req *types.RoleDelete) (resp types.Response) {

	c := l.svcCtx
	ids := req.Ids
	db := c.Db
	role := model.Role{}
	e, err := casbin.NewEnforcer("")
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	if len(ids) == 0 {
		id := req.Id
		if id == "" {
			resp.Error("id不能为空！", nil)
			return
		}

		tx := db.Where("id = ?", id).First(&role)

		if util.IsDbErr(tx) != nil {
			resp.Error(tx.Error.Error(), nil)
			return
		}

		// 删除对应的角色关系
		e.DeleteRole(strconv.Itoa(role.Id))
		if err = db.Where("id = ?", id).Delete(&role).Error; err != nil {
			resp.Error("删除失败！", err.Error())
			return
		}
	} else {
		if err = db.Where("id IN (?)", ids).Delete(&role).Error; err != nil {
			resp.Error("删除失败！", nil)
			return
		}

		for _, v := range ids {
			e.DeleteRole(v)
		}

	}

	resp.Success("删除成功！", nil)
	return
}
