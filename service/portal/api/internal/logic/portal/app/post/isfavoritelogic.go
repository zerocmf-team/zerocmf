package post

import (
	"context"
	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"
	"zerocmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteLogic {
	return &IsFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFavoriteLogic) IsFavorite(req *types.OneReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
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
