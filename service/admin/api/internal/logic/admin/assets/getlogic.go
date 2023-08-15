package assets

import (
	"context"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/admin/model"

	"zerocmf/service/admin/api/internal/svc"
	"zerocmf/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.AssetsReq) (resp *types.Response) {
	resp = new(types.Response)
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := c.Request
	userId, _ := c.Get("userId")
	query := []string{"user_id = ? AND status = ?"}
	queryArgs := []interface{}{userId, "1"}
	paramType := req.Type
	query = append(query, "type = ?")
	queryArgs = append(queryArgs, paramType)
	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := data.NewPaginate(r).Default()

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	data, err := new(model.Assets).Get(db, current, pageSize, queryStr, queryArgs)

	if err != nil {
		resp.Error("系统出错", nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
