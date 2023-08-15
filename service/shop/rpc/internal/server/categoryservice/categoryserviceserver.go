// Code generated by goctl. DO NOT EDIT.
// Source: shop.proto

package server

import (
	"context"

	"zerocmf/service/shop/rpc/internal/logic/categoryservice"
	"zerocmf/service/shop/rpc/internal/svc"
	"zerocmf/service/shop/rpc/pb/shop"
)

type CategoryServiceServer struct {
	svcCtx *svc.ServiceContext
	shop.UnimplementedCategoryServiceServer
}

func NewCategoryServiceServer(svcCtx *svc.ServiceContext) *CategoryServiceServer {
	return &CategoryServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CategoryServiceServer) CategoryGet(ctx context.Context, in *shop.CategoryGetReq) (*shop.CategoryListResp, error) {
	l := categoryservicelogic.NewCategoryGetLogic(ctx, s.svcCtx)
	return l.CategoryGet(in)
}

func (s *CategoryServiceServer) CategoryTree(ctx context.Context, in *shop.CategoryTreeReq) (*shop.CategoryTreeListResp, error) {
	l := categoryservicelogic.NewCategoryTreeLogic(ctx, s.svcCtx)
	return l.CategoryTree(in)
}

func (s *CategoryServiceServer) CategoryShow(ctx context.Context, in *shop.CategoryShowReq) (*shop.CategoryResp, error) {
	l := categoryservicelogic.NewCategoryShowLogic(ctx, s.svcCtx)
	return l.CategoryShow(in)
}

func (s *CategoryServiceServer) CategorySave(ctx context.Context, in *shop.CategorySaveReq) (*shop.CategoryResp, error) {
	l := categoryservicelogic.NewCategorySaveLogic(ctx, s.svcCtx)
	return l.CategorySave(in)
}

func (s *CategoryServiceServer) CategoryDel(ctx context.Context, in *shop.CategoryDelReq) (*shop.CategoryResp, error) {
	l := categoryservicelogic.NewCategoryDelLogic(ctx, s.svcCtx)
	return l.CategoryDel(in)
}
