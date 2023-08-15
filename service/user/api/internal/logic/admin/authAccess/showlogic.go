package authAccess

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id
	if id == "" {
		resp.Error("角色id不能为空！", nil)
	}

	adapter := c.Config.Database.NewConf(siteId.(int64))
	e, err := adapter.NewEnforcer()
	if err != nil {
		return
	}
	role := model.Role{}
	tx := db.Where("id = ? AND status = 1", id).First(&role)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			resp.Error("该角色不存在或已删除！", nil)
			return
		}
		resp.Error(tx.Error.Error(), nil)
		return
	}

	id = strconv.Itoa(role.Id)

	// 获取全部角色策略
	roles := e.GetFilteredPolicy(0, id)
	result := make([]string, 0)
	for _, v := range roles {
		if len(v) > 1 {
			result = append(result, v[1])
		}
	}
	resp.Success("获取成功！", result)
	return
}
