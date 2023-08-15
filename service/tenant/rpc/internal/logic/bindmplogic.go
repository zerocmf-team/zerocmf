package logic

import (
	"context"
	"github.com/zerocmf/wechatEasySdk/wxopen"
	"strings"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindMpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindMpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindMpLogic {
	return &BindMpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindMpLogic) BindMp(in *tenant.BindMpReq) (reply *tenant.BindMpReply, err error) {
	reply = new(tenant.BindMpReply)
	c := l.svcCtx
	db := c.Db
	query := []string{"auth_app_id = ?"}
	authAppId := in.AuthorizerAppid
	queryArgs := []interface{}{authAppId}
	queryStr := strings.Join(query, " AND ")
	auth := model.SiteMpAuth{}
	tx := db.Where(queryStr, queryArgs...).Order("id desc").First(&auth)
	if util.IsDbErr(tx) != nil {
		reply.Status = false
		reply.Msg = tx.Error.Error()
		return
	}

	if err != nil {
		reply.Status = false
		reply.Msg = tx.Error.Error()
		return
	}

	auth.SiteId = in.GetSiteId()
	auth.Type = in.GetType()
	auth.AuthAppId = in.GetAuthorizerAppid()
	auth.AppAuthToken = in.GetAuthorizerAccessToken()
	auth.AppRefreshToken = in.GetAuthorizerRefreshToken()
	auth.ExpiresIn = in.GetExpiresIn()

	if auth.Id == 0 {
		auth.CreatedAt = time.Now().Unix()
		tx = db.Create(&auth)
		if util.IsDbErr(tx) != nil {
			reply.Status = false
			reply.Msg = tx.Error.Error()
			return
		}
	} else {
		if authAppId != auth.AuthAppId {
			reply.Status = false
			reply.Msg = "重新授权的账号与当前绑定的账号不一致"
			return
		}
		auth.UpdatedAt = time.Now().Unix()
		tx = db.Updates(&auth)
		if util.IsDbErr(tx) != nil {
			reply.Status = false
			reply.Msg = tx.Error.Error()
			return
		}
	}

	bizContent := wxopen.ModifyDomain{
		AuthorizerAccessToken: auth.AppAuthToken,
		Action:                "set",
		RequestDomain: []string{
			"http://www.zerocms.cn",
		},
		WSRequestDomain: []string{
			"http://www.zerocms.cn",
		},
		UploadDomain: []string{
			"http://www.zerocms.cn",
		},
		DownloadDomain: []string{
			"http://www.zerocms.cn",
		},
	}
	new(wxopen.Wxa).ModifyDomain(bizContent)
	//var domainResult wxopen.ModifyDomainResult
	//domainResult, err = new(wxopen.Wxa).ModifyDomain(bizContent)
	//if err != nil {
	//	return
	//}
	//if domainResult.ErrCode > 0 {
	//	fmt.Println("domainResult", domainResult)
	//	reply.Status = false
	//	reply.Msg = domainResult.ErrMsg
	//	return
	//}

	bizContent1 := wxopen.SetWebViewDomain{
		AuthorizerAccessToken: auth.AppAuthToken,
		Action:                "set",
		WebViewDomain: []string{
			"http://www.zerocms.cn",
		},
	}
	new(wxopen.Wxa).SetWebViewDomain(bizContent1)
	//var result2 wxopen.WebViewDomainResult
	//result2, err = new(wxopen.Wxa).SetWebViewDomain(bizContent1)
	//if err != nil {
	//	return
	//}
	//if result2.ErrCode > 0 {
	//	reply.Status = false
	//	reply.Msg = domainResult.ErrMsg
	//	return
	//}
	reply.Status = true
	reply.Msg = "绑定成功！"
	return
}
