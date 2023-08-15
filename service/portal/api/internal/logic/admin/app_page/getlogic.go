package app_page

import (
	"context"
	"net/http"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetLogic {
	ctx := header.Context()
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.AppPageListReq) (resp types.Response) {
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	db := c.Config.Database.ManualDb(siteId.(int64))
	r := l.header

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	query := []string{"app_id = ?", "delete_at = ?"}
	queryArgs := []interface{}{req.AppId, 0}

	typ := req.Type
	if typ != "" {
		query = append(query, "type = ?")
		queryArgs = append(queryArgs, typ)
	}

	isPublic := req.IsPublic
	if isPublic != nil {
		query = append(query, "is_public = ?")
		queryArgs = append(queryArgs, isPublic)
	}

	if req.Name != nil {
		query = append(query, "name like ?")
		queryArgs = append(queryArgs, "%"+*req.Name+"%")
	}

	if req.Status != nil {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, req.Status)
	}

	queryStr := strings.Join(query, " AND ")
	appPage := new(model.AppPage)
	paginate := req.Paginate
	var result interface{}
	if paginate == "no" {
		var list []model.AppPage
		list, err = appPage.List(db, queryStr, queryArgs)
		if err != nil {
			resp.Error("获取失败！", err.Error())
			return
		}
		result = list
	} else {
		var pageData data.Paginate
		pageData, err = appPage.Index(db, current, pageSize, queryStr, queryArgs)

		if err != nil {
			resp.Error("获取失败！", err.Error())
			return
		}
		result = pageData
	}

	resp.Success("获取成功！", result)
	return
}
