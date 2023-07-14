package comment

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/user/rpc/userclient"
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
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(string))
	r := c.Request
	userRpc := c.UserRpc

	topicId := req.Id
	typ := req.Type

	query := "topic_id = ? AND topic_type = ?"
	queryArgs := []interface{}{topicId, typ}

	userId, _ := c.Get("userId")
	userIdInt, _ := strconv.Atoi(userId.(string))

	if userId != "" {

		_, err := userRpc.Get(context.Background(), &userclient.UserRequest{
			UserId: userId.(string),
			SiteId: siteId.(string),
		})

		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}

	}

	current, pageSize, err := data.NewPaginate(r).Default()
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
