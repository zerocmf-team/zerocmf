package list

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/portal/model"
	"strings"

	"zerocmf/service/portal/api/internal/svc"
	"zerocmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.ArticleSearchReq) (resp types.Response) {
	c := l.svcCtx
	r := c.Request
	db := c.Db
	keywords := req.Keywords

	if strings.TrimSpace(keywords) == "" {
		resp.Error("关键字不能为空", nil)
		return
	}

	query := []string{"p.post_type = ?","p.delete_at = ?","(p.post_title like ? or p.post_keywords like ? or p.post_excerpt like ? or p.post_content like ?)"}
	queryArgs := []interface{}{1,0,"%"+keywords+"%","%"+keywords+"%","%"+keywords+"%","%"+keywords+"%"}

	queryStr := strings.Join(query, " AND ")

	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	post := model.PortalPost{}
	res, err := post.IndexByCategory(db, current, pageSize, queryStr, queryArgs, nil)
	if err != nil {
		resp.Error("获取失败！", nil)
		return
	}

	resp.Success("获取成功！", res)
	return
}
