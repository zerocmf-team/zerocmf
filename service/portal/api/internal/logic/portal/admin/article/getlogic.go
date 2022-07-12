package article

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"
	"strings"
	"time"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.ArticleGetReq) (resp types.Response) {
	c := l.svcCtx
	r := c.Request
	db := c.Db

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

	startTime := req.StartTime
	endTime := req.EndTime

	if startTime != "" && endTime != "" {
		startTimeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
		if err != nil {
			resp.Error("起始时间非法！", err.Error())
			return
		}

		endTimeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
		if err != nil {
			resp.Error("结束时间非法！", err.Error())
		}

		query = append(query, "((p.publish_at BETWEEN ? AND ?) OR (p.update_at BETWEEN ? AND ?))")
		queryArgs = append(queryArgs, startTimeStamp, endTimeStamp, startTimeStamp, endTimeStamp)
	}
	queryStr := strings.Join(query, " AND ")
	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	post := model.PortalPost{}
	data, err := post.IndexByCategory(db, current, pageSize, queryStr, queryArgs, nil)
	if err != nil {
		resp.Error("获取失败！", nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
