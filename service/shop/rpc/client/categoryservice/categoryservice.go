// Code generated by goctl. DO NOT EDIT.
// Source: shop.proto

package categoryservice

import (
	"context"

	"zerocmf/service/shop/rpc/pb/shop"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Attributes           = shop.Attributes
	AttributesItem       = shop.AttributesItem
	CategoryDelReq       = shop.CategoryDelReq
	CategoryGetReq       = shop.CategoryGetReq
	CategoryListResp     = shop.CategoryListResp
	CategoryResp         = shop.CategoryResp
	CategorySaveReq      = shop.CategorySaveReq
	CategoryShowReq      = shop.CategoryShowReq
	CategoryTreeData     = shop.CategoryTreeData
	CategoryTreeListResp = shop.CategoryTreeListResp
	CategoryTreeReq      = shop.CategoryTreeReq
	MigrateReply         = shop.MigrateReply
	MigrateReq           = shop.MigrateReq
	ProductAttrKeyReq    = shop.ProductAttrKeyReq
	ProductAttrKeyResp   = shop.ProductAttrKeyResp
	ProductAttrValReq    = shop.ProductAttrValReq
	ProductAttrValResp   = shop.ProductAttrValResp
	ProductGetReq        = shop.ProductGetReq
	ProductListResp      = shop.ProductListResp
	ProductResp          = shop.ProductResp
	ProductSaveReq       = shop.ProductSaveReq
	ProductSaveResp      = shop.ProductSaveResp
	ProductShowReq       = shop.ProductShowReq
	ProductSku           = shop.ProductSku

	CategoryService interface {
		CategoryGet(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		CategoryTree(ctx context.Context, in *CategoryTreeReq, opts ...grpc.CallOption) (*CategoryTreeListResp, error)
		CategoryShow(ctx context.Context, in *CategoryShowReq, opts ...grpc.CallOption) (*CategoryResp, error)
		CategorySave(ctx context.Context, in *CategorySaveReq, opts ...grpc.CallOption) (*CategoryResp, error)
		CategoryDel(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*CategoryResp, error)
	}

	defaultCategoryService struct {
		cli zrpc.Client
	}
)

func NewCategoryService(cli zrpc.Client) CategoryService {
	return &defaultCategoryService{
		cli: cli,
	}
}

func (m *defaultCategoryService) CategoryGet(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := shop.NewCategoryServiceClient(m.cli.Conn())
	return client.CategoryGet(ctx, in, opts...)
}

func (m *defaultCategoryService) CategoryTree(ctx context.Context, in *CategoryTreeReq, opts ...grpc.CallOption) (*CategoryTreeListResp, error) {
	client := shop.NewCategoryServiceClient(m.cli.Conn())
	return client.CategoryTree(ctx, in, opts...)
}

func (m *defaultCategoryService) CategoryShow(ctx context.Context, in *CategoryShowReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	client := shop.NewCategoryServiceClient(m.cli.Conn())
	return client.CategoryShow(ctx, in, opts...)
}

func (m *defaultCategoryService) CategorySave(ctx context.Context, in *CategorySaveReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	client := shop.NewCategoryServiceClient(m.cli.Conn())
	return client.CategorySave(ctx, in, opts...)
}

func (m *defaultCategoryService) CategoryDel(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	client := shop.NewCategoryServiceClient(m.cli.Conn())
	return client.CategoryDel(ctx, in, opts...)
}