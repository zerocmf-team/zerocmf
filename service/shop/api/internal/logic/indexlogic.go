package logic

import (
	"context"
	"net/http"
	"zerocmf/common/bootstrap/data"

	"github.com/zeromicro/go-zero/core/logx"
	"zerocmf/service/shop/api/internal/svc"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(header *http.Request, svcCtx *svc.ServiceContext) *IndexLogic {
	ctx := header.Context()
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp data.Rest) {
	statusCode := 201
	resp.StatusCode = &statusCode
	resp.Success("获取成功！", data.H{
		"author":  "zerocmf",
		"message": "商城系统",
		"name":    "shop",
	})
	return
}
