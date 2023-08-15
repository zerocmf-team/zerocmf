package login

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(header *http.Request, svcCtx *svc.ServiceContext) *LoginLogic {
	ctx := header.Context()
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp data.Rest) {

	c := l.svcCtx
	redis := c.Redis
	siteId, _ := c.Get("siteId")

	componentAccessToken, exist := c.Get("componentAccessToken")
	if !exist {
		resp.Error("accessToken不存在！", nil)
		return
	}

	var err error

	bizContent := wxopen.Code2Session{
		ComponentAppid:       c.Config.Wechat.WxOpen.ComponentAppId,
		ComponentAccessToken: componentAccessToken.(string),
	}
	err = copier.Copy(&bizContent, &req)
	if err != nil {
		resp.Error("请求错误", err.Error())
		return
	}

	var sessionResult wxopen.Code2SessionResult
	sessionResult, err = new(wxopen.Component).Code2session(bizContent)
	if err != nil {
		resp.Error("请求错误", err.Error())
		return
	}

	if sessionResult.ErrCode > 0 {
		redis.Del("componentAccessToken")
		resp.Error("登录失败！请重试", sessionResult)
		return
	}

	//openId := sessionResult.Openid

	// 根据openId创建唯一用户

	var result = struct {
		OpenId string
		SiteId int64
	}{
		OpenId: sessionResult.Openid,
		SiteId: siteId.(int64),
	}

	resp.Success("登录成功！", result)
	return
}
