package comment

import (
	"context"
	"net/http"
	"strconv"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/model"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.PostCommentGetReq) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := l.header
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
			SiteId: siteId.(int64),
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
