package admin

import (
	"context"
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"net/http"
	"strconv"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/tenant/rpc/tenantclient"
	"zerocmf/service/tenant/rpc/types/tenant"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindMpLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewBindMpLogic(header *http.Request, svcCtx *svc.ServiceContext) *BindMpLogic {
	ctx := header.Context()
	return &BindMpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *BindMpLogic) BindMp(req *types.BindMpReq) (resp data.Rest) {
	c := l.svcCtx

	componentAccessToken, exist := c.Get("componentAccessToken")
	if !exist {
		resp.Error("accessToken不存在！", nil)
		return
	}

	siteId, _ := c.Get("siteId")

	typ := req.Typ

	// 查询授权
	bizContent := wxopen.QueryAuth{
		ComponentAccessToken: componentAccessToken.(string),
		ComponentAppId:       c.Config.Wechat.WxOpen.ComponentAppId,
		AuthorizationCode:    req.AuthCode,
	}
	result, err := new(wxopen.Component).QueryAuth(bizContent)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	if result.ErrCode > 0 {
		resp.Error("请求失败！", result)
		return
	}

	var reply *tenantclient.BindMpReply
	reply, err = c.TenantRpc.BindMp(l.ctx, &tenant.BindMpReq{
		SiteId:                 siteId.(int64),
		Type:                   typ,
		AuthorizerAppid:        result.AuthorizationInfo.AuthorizerAppid,
		AuthorizerAccessToken:  result.AuthorizationInfo.AuthorizerAccessToken,
		AuthorizerRefreshToken: result.AuthorizationInfo.AuthorizerRefreshToken,
		ExpiresIn:              strconv.Itoa(result.AuthorizationInfo.ExpiresIn),
	})
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}
	if reply.Status {
		resp.Success(reply.Msg, nil)
		return
	}
	resp.Error(reply.Msg, nil)
	return
}
