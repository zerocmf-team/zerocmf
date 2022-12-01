package account

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/user/model"
	"strings"

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

func (l *GetLogic) Get(req *types.ListReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	r := c.Request

	query := []string{"user_type = ?"}
	queryArgs := []interface{}{"1"}

	userLogin := req.UserLogin
	if userLogin != "" {
		query = append(query, "user_login LIKE ?")
		queryArgs = append(queryArgs, "%"+userLogin+"%")
	}

	userNickname := req.UserNickname
	if userNickname != "" {
		query = append(query, "user_nickname like ?")
		queryArgs = append(queryArgs, "%"+userNickname+"%")
	}

	userEmail := req.UserEmail
	if userEmail != "" {
		query = append(query, "user_email like ?")
		queryArgs = append(queryArgs, "%"+userEmail+"%")
	}

	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := data.NewPaginate(r).Default()

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	result, err := new(model.User).Paginate(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		resp.Error("获取失败："+err.Error(), nil)
		return
	}
	resp.Success("获取成功！", result)

	return
}
