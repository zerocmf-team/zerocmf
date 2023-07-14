package role

import (
	"context"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/user/model"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

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

func (l *GetLogic) Get(req *types.RoleGet) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	r := c.Request

	var query []string
	var queryArgs []interface{}

	status := req.Status
	if status != "" {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, status)
	}
	// 名称模糊搜索
	name := req.Name
	if name != "" {
		query = append(query, "name LIKE ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	var (
		queryStr string
		current  int
		pageSize int
		err      error
	)
	queryStr = strings.Join(query, " AND ")

	current, pageSize, err = data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	if pageSize > 0 {
		var result data.Paginate
		result, err = new(model.Role).Paginate(db, current, pageSize, queryStr, queryArgs)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		resp.Success("获取成功！", result)
	} else {
		var result []model.Role
		result, err = new(model.Role).List(db, queryStr, queryArgs)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		resp.Success("获取成功！", result)
	}

	return
}
