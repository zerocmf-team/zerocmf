package logic

import (
	"context"
	"errors"
	"strings"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/tenant/model"

	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowMpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowMpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowMpLogic {
	return &ShowMpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowMpLogic) ShowMp(in *tenant.ShowMpData) (*tenant.ShowMpData, error) {

	c := l.svcCtx
	db := c.Db

	var query []string
	var queryArgs []interface{}

	siteId := in.GetSiteId()
	if siteId > 0 {
		query = append(query, "site_id = ?")
		queryArgs = append(queryArgs, siteId)
	}

	typ := in.GetType()
	if typ != "" {
		query = append(query, "type = ?")
		queryArgs = append(queryArgs, typ)
	}

	appId := in.GetAppId()
	if appId != "" {
		query = append(query, "auth_app_id = ?")
		queryArgs = append(queryArgs, appId)
	}

	if len(query) == 0 {
		return &tenant.ShowMpData{}, errors.New("参数不能为空！")
	}

	queryStr := strings.Join(query, " AND ")

	mpAuth := model.SiteMpAuth{}
	tx := db.Where(queryStr, queryArgs...).First(&mpAuth)
	if util.IsDbErr(tx) != nil {
		return &tenant.ShowMpData{}, errors.New(tx.Error.Error())
	}

	return &tenant.ShowMpData{
		SiteId: mpAuth.SiteId,
		Type:   mpAuth.Type,
		AppId:  mpAuth.AuthAppId,
	}, nil
}
