package category

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/categoryservice"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

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

func (l *GetLogic) Get(req *types.CategoryGetReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx

	current, pageSize, err := data.PaginateQuery(l.header)
	if err != nil {
		resp.Error("系统错误", nil)
		return
	}

	categoryClient := categoryservice.NewCategoryService(c.Client)
	rpcReq := categoryservice.CategoryGetReq{}
	copier.Copy(&rpcReq, &req)
	rpcReq.Current = current
	rpcReq.PageSize = pageSize
	var category *categoryservice.CategoryListResp
	category, err = categoryClient.CategoryGet(ctx, &rpcReq)
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	var json interface{} = category

	if pageSize > 0 {
		json = data.Paginate{
			Current:  current,
			PageSize: pageSize,
			Data:     category.GetData(),
			Total:    category.GetTotal(),
		}
	}

	resp.Success("获取成功", json)
	return
}
