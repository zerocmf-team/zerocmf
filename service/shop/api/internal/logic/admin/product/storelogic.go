package product

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"net/http"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/shop/rpc/client/productservice"

	"zerocmf/service/shop/api/internal/svc"
	"zerocmf/service/shop/api/internal/types"

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

func (l *StoreLogic) Store(req *types.ProductSaveReq) (resp data.Rest) {
	ctx := l.ctx
	c := l.svcCtx
	siteId, _ := c.Get("siteId")
	client := productservice.NewProductService(c.Client)
	rpcReq := productservice.ProductSaveReq{}
	rpcReq.SiteId = siteId.(int64)
	err := copier.Copy(&rpcReq, &req)
	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	attributes := req.Attributes
	attrJson, aErr := json.Marshal(attributes)
	if aErr != nil {
		resp.Error("系统出错了", err.Error())
		return
	}
	rpcReq.Attributes = string(attrJson)

	var rpcResp *productservice.ProductSaveResp
	rpcResp, err = client.ProductSave(ctx, &rpcReq)
	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	msg := "添加成功！"
	if req.ProductId > 0 {
		msg = "修改成功！"
	}

	resp.Success(msg, rpcResp)
	return
}
