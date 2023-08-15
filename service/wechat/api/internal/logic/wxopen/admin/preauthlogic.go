package admin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"zerocmf/common/bootstrap/data"
	"zerocmf/common/bootstrap/util"

	"github.com/zerocmf/wechatEasySdk/wxopen"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreAuthLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewPreAuthLogic(header *http.Request, svcCtx *svc.ServiceContext) *PreAuthLogic {
	ctx := header.Context()
	return &PreAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *PreAuthLogic) PreAuth(req *types.PreAuthReq) (resp data.Rest) {

	c := l.svcCtx
	r := l.header

	siteId, _ := c.Get("siteId")

	componentAccessToken, exist := c.Get("componentAccessToken")
	if !exist {
		resp.Error("accessToken不存在！", nil)
		return
	}

	var err error

	componentAppId := c.Config.Wechat.WxOpen.ComponentAppId
	accessToken := componentAccessToken.(string)

	bizContent := wxopen.PreAuthCode{
		ComponentAppid:       componentAppId,
		ComponentAccessToken: accessToken,
	}

	var result wxopen.PreAuthCodeResult
	result, err = new(wxopen.Component).PreAuthCode(bizContent)
	if err != nil {
		return
	}

	if result.ErrCode == 0 {
		host := req.Redirect
		if host == "" {
			host = util.Host(r)
		}

		stateMap := make(map[string]interface{}, 0)
		stateMap["siteId"] = siteId
		stateMap["type"] = "wxapp"
		var bytes []byte
		bytes, err = json.Marshal(stateMap)
		if err != nil {
			fmt.Println("err", err)
		}
		state := base64.StdEncoding.EncodeToString(bytes)
		query := url.Values{}
		query.Add("state", state)
		queryParams := query.Encode()

		preAuthCode := result.PreAuthCode

		siteIdInt := siteId.(int64)

		redirectUrl := url.QueryEscape(host + "/" + strconv.FormatInt(siteIdInt, 10) + c.Config.Wechat.WxOpen.RedirectUrl + "?" + queryParams)
		url := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + componentAppId + "&pre_auth_code=" + preAuthCode + "&auth_type=2&redirect_uri=" + redirectUrl
		resp.Success("获取成功！", url)
	}

	return
}
