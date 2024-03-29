package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/user/common"
	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"
)

type ValidationJwtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidationJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidationJwtLogic {
	return &ValidationJwtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidationJwtLogic) ValidationJwt(in *user.OauthRequest) (*user.OauthReply, error) {
	token := in.GetToken()
	tenantId := in.GetTenantId()
	c := l.svcCtx
	conf := c.Config
	inConf := common.Config{}
	copier.Copy(&inConf, &conf)
	oauth := common.NewServer(inConf, tenantId)
	srv := oauth.Srv
	ti, err := srv.Manager.LoadAccessToken(context.Background(), token)
	if err != nil {

		return &user.OauthReply{}, err
	}
	return &user.OauthReply{
		UserId: ti.GetUserID(),
	}, nil
}
