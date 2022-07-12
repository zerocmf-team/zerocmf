package category

import (
	"context"
	"zerocmf/service/portal/model"
	"strings"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionsLogic {
	return &OptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionsLogic) Options() (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	category := model.PortalCategory{}
	var query = []string{"delete_at  = ?"}
	var queryArgs = []interface{}{0}
	queryStr := strings.Join(query, " AND ")

	data, err := category.ListWithOptions(db, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
