package assets

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/admin/model"
	"strings"

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

func (l *GetLogic) Get(req types.AssetsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{}
	c := l.svcCtx
	db := c.Db
	r := c.Request
	userId, _ := c.Get("userId")
	query := []string{"user_id = ? AND status = ?"}
	queryArgs := []interface{}{userId, "1"}
	paramType := req.Type
	query = append(query, "type = ?")
	queryArgs = append(queryArgs, paramType)
	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := new(data.Paginate).Default(r)

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
