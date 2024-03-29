package article

import (
	"context"
	"net/http"
	"strings"
	"time"
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

func (l *GetLogic) Get(req *types.ArticleGetReq) (resp types.Response) {
	c := l.svcCtx
	r := l.header
	siteId, _ := c.Get("siteId")
	db := c.NewDb(siteId.(int64))

	query := []string{"p.delete_at = ?"}
	queryArgs := []interface{}{0}

	title := req.Title
	if title != "" {
		query = append(query, "p.post_title like ?")
		queryArgs = append(queryArgs, "%"+title+"%")
	}

	postType := req.PostType
	if postType == "2" {
		postType = "2"
	} else {
		postType = "1"
	}

	query = append(query, "p.post_type = ?")
	queryArgs = append(queryArgs, postType)

	Category := req.Category
	if Category != nil {
		query = append(query, "pc.id = ?")
		queryArgs = append(queryArgs, Category)
	}

	postStatus := req.PostStatus
	if postStatus != nil {
		query = append(query, "p.post_status = ?")
		queryArgs = append(queryArgs, postStatus)
	}

	startTime := req.StartTime
	endTime := req.EndTime

	var (
		startTimeStamp time.Time
		endTimeStamp   time.Time
		err            error
	)
	if startTime != "" && endTime != "" {
		startTimeStamp, err = time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
		if err != nil {
			resp.Error("起始时间非法！", err.Error())
			return
		}

		endTimeStamp, err = time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
		if err != nil {
			resp.Error("结束时间非法！", err.Error())
		}

		query = append(query, "((p.publish_at BETWEEN ? AND ?) OR (p.update_at BETWEEN ? AND ?))")
		queryArgs = append(queryArgs, startTimeStamp, endTimeStamp, startTimeStamp, endTimeStamp)
	}
	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := data.NewPaginate(r).Default()
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	post := model.PortalPost{}
	var pageData data.Paginate
	pageData, err = post.ListByCategory(db, current, pageSize, queryStr, queryArgs, nil)
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	resp.Success("获取成功！", pageData)
	return
}
