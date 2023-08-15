package fastRegisterWeApp

import (
	"context"
	"github.com/jinzhu/copier"
	wxData "github.com/zerocmf/wechatEasySdk/data"
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FastRegisterWeAppLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewFastRegisterWeAppLogic(header *http.Request, svcCtx *svc.ServiceContext) *FastRegisterWeAppLogic {
	ctx := header.Context()
	return &FastRegisterWeAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *FastRegisterWeAppLogic) FastRegisterWeApp(req *types.FastRegisterWeAppReq) (resp data.Rest) {

	c := l.svcCtx

	componentAccessToken, exist := c.Get("componentAccessToken")
	if !exist {
		resp.Error("accessToken不存在！", nil)
		return
	}

	if req.Pwd != "codecloud2021" {
		resp.Error("安全秘钥验证失败！", nil)
		return
	}

	if req.Name == "" {
		resp.Error("企业名不能为空", nil)
		return
	}

	if req.CodeType == 0 || req.CodeType > 3 {
		resp.Error("企业代码类型错误或为空", nil)
		return
	}

	if req.LegalPersonaWechat == "" {
		resp.Error("法人微信不能为空", nil)
		return
	}

	if req.LegalPersonaName == "" {
		resp.Error("法人姓名不能为空", nil)
		return
	}

	if req.ComponentPhone == "" {
		req.ComponentPhone = "17177723588"
	}

	bizContent := wxopen.FastRegisterWeapp{
		ComponentAccessToken: componentAccessToken.(string),
	}
	err := copier.Copy(&bizContent, &req)
	if err != nil {
		resp.Error("请求失败！", nil)
		return
	}

	var wxResp = wxData.Response{}
	wxResp, err = new(wxopen.Component).FastRegisterWeapp(bizContent)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	if wxResp.ErrCode != 0 {
		resp.Error(wxResp.ErrMsg, nil)
		return
	}
	resp.Success("发起成功，请尽快处理", nil)
	return
}
