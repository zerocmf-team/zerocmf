package post

import (
	"context"
	"net/http"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewIsFavoriteLogic(header *http.Request, svcCtx *svc.ServiceContext) *IsFavoriteLogic {
	ctx := header.Context()
	return &IsFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *IsFavoriteLogic) IsFavorite(req *types.OneReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	id := req.Id
	if id == 0 {
		resp.Error("id不能为空", nil)
		return
	}
	userId, _ := c.Get("userId")
	favoritesPost := new(model.PostFavoritesPost)
	prefix := c.Config.Database.Prefix
	favoritesPost.Table = prefix + "post_favorites_post"
	err := favoritesPost.IsLike(db, id, userId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", favoritesPost.Status)
	return
}
