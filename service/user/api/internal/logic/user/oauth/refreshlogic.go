package oauth

import (
	"context"
	"zerocmf/service/user/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/oauth2"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp types.Response) {
	refreshToken := req.RefreshToken
	if refreshToken == "" {
		resp.Error("refresh_token不能为空！", nil)
		return
	}
	token := &oauth2.Token{RefreshToken: refreshToken}

	c := l.svcCtx

	conf := c.Config
	inConf := common.Config{}
	copier.Copy(&inConf, &conf)
	oauthConfig := common.NewConf(inConf,"")
	tkr := oauthConfig.TokenSource(context.Background(), token)
	tk, err := tkr.Token()

	if err != nil {
		resp.Error("获取失败："+err.Error(), gin.H{
			"error":             "invalid_client",
			"error_description": "Client authentication failed",
		})
		return
	}
	resp.Success("获取成功！", tk)
	return
}
