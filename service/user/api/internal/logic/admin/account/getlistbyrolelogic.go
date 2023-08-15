package account

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByRoleLogic {
	return &GetListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListByRoleLogic) GetListByRole(req *types.ListByRoleReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	dbConf := c.Config.Database.NewConf(siteId.(int64))
	db := dbConf.ManualDb(siteId.(int64))
	//pageSize := req.PageSize

	roleIds := strings.Join(req.RoleIds, ",")

	var result []model.User

	prefix := dbConf.Prefix
	tx := db.Select("u.*").Table("casbin_rule rule").Joins("INNER JOIN "+prefix+"user u ON rule.v0 = u.id").
		Where("rule.v1 in (?)", roleIds).Scan(&result)

	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		resp.Error("获取失败！", tx.Error.Error())
		return
	}

	resp.Success("获取成功", result)
	return
}
