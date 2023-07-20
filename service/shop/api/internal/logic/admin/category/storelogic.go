package category

import (
	"context"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"
	"zerocmf/service/shop/rpc/client/categoryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(header *http.Request, svcCtx *svc.ServiceContext) *StoreLogic {
	ctx := header.Context()
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.CategorySaveReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	//siteId, _ := c.Get("siteId")
	//siteStr := siteId.(string)
	categoryClient := categoryservice.NewCategoryService(c.Client)
	categoryReq := categoryservice.CategorySaveReq{}
	err := copier.Copy(&categoryReq, &req)
	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	var listOrder float64 = 10000
	if req.ListOrder != nil {
		listOrder = *req.ListOrder
	}

	categoryReq.ListOrder = listOrder

	status := 1
	if req.Status != nil {
		status = *req.Status
	}
	categoryReq.Status = int32(status)

	var categoryResp *categoryservice.CategoryResp
	categoryResp, err = categoryClient.CategorySave(ctx, &categoryReq)

	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	msg := "添加成功！"
	if req.Id > 0 {
		msg = "修改成功！"
	}

	resp.Success(msg, categoryResp)
	return
}
