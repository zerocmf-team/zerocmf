package comment

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/user/rpc/user"
	"strconv"

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

func (l *GetLogic) Get(req *types.PostCommentGetReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	r := c.Request
	userRpc := c.UserRpc

	topicId := req.Id
	typ := req.Type

	query := "topic_id = ? AND topic_type = ?"
	queryArgs := []interface{}{topicId, typ}

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	if userId != "" {
		tenant, exist := db.Get("tenantId")
		tenantId := ""
		if exist {
			tenantId = tenant.(string)
		}

		_, err := userRpc.Get(context.Background(), &user.UserRequest{
			UserId:   int64(userIdInt),
			TenantId: tenantId,
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

	}

	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	data, err := new(model.Comment).Paginate(db, current, pageSize, query, queryArgs, userIdInt)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
