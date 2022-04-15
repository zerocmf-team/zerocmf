package oauth

import (
	"context"
	"gincmf/service/user/common"
	"github.com/jinzhu/copier"

	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenRequestLogic {
	return &TokenRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenRequestLogic) TokenRequest() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	c := l.svcCtx
	r := c.Request
	w := c.ResponseWriter

	conf := c.Config
	inConf := common.Config{}
	copier.Copy(&inConf, &conf)

	oauth := common.NewServer(inConf, "")
	defer oauth.Store.Close()
	srv := oauth.Srv

	err = srv.HandleTokenRequest(w, r)
	if err != nil {
		return
	}
	return
}
