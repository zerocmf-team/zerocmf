package form

import (
	"context"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.FormListReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	r := c.Request

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	query := []string{"delete_at = ?"}
	queryArgs := []interface{}{0}
	queryStr := strings.Join(query, " AND ")

	form := model.Form{}

	res, resErr := form.List(db, current, pageSize, queryStr, queryArgs, true)

	if resErr != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功!", res)
	return
}
