package post

import (
	"context"
	"net/http"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLikeLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewIsLikeLogic(header *http.Request, svcCtx *svc.ServiceContext) *IsLikeLogic {
	ctx := header.Context()
	return &IsLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *IsLikeLogic) IsLike(req *types.OneReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))

	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}

	userId, _ := c.Get("userId")

	postLikePost := new(model.PostLikePost)
	prefix := c.Config.Database.Prefix
	postLikePost.Table = prefix + "post_like_post"

	err := postLikePost.IsLike(db, id, userId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", postLikePost.Status)
	return
}
