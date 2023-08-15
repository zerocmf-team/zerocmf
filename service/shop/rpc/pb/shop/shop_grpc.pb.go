// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pb/shop.proto

package shop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ShopService_AutoMigrate_FullMethodName = "/shop.ShopService/AutoMigrate"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	AutoMigrate(ctx context.Context, in *MigrateReq, opts ...grpc.CallOption) (*MigrateReply, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) AutoMigrate(ctx context.Context, in *MigrateReq, opts ...grpc.CallOption) (*MigrateReply, error) {
	out := new(MigrateReply)
	err := c.cc.Invoke(ctx, ShopService_AutoMigrate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	AutoMigrate(context.Context, *MigrateReq) (*MigrateReply, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) AutoMigrate(context.Context, *MigrateReq) (*MigrateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutoMigrate not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_AutoMigrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AutoMigrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_AutoMigrate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AutoMigrate(ctx, req.(*MigrateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AutoMigrate",
			Handler:    _ShopService_AutoMigrate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/shop.proto",
}

const (
	CategoryService_CategoryGet_FullMethodName  = "/shop.CategoryService/CategoryGet"
	CategoryService_CategoryTree_FullMethodName = "/shop.CategoryService/CategoryTree"
	CategoryService_CategoryShow_FullMethodName = "/shop.CategoryService/CategoryShow"
	CategoryService_CategorySave_FullMethodName = "/shop.CategoryService/CategorySave"
	CategoryService_CategoryDel_FullMethodName  = "/shop.CategoryService/CategoryDel"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CategoryGet(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*CategoryListResp, error)
	CategoryTree(ctx context.Context, in *CategoryTreeReq, opts ...grpc.CallOption) (*CategoryTreeListResp, error)
	CategoryShow(ctx context.Context, in *CategoryShowReq, opts ...grpc.CallOption) (*CategoryResp, error)
	CategorySave(ctx context.Context, in *CategorySaveReq, opts ...grpc.CallOption) (*CategoryResp, error)
	CategoryDel(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*CategoryResp, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CategoryGet(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	out := new(CategoryListResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoryGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoryTree(ctx context.Context, in *CategoryTreeReq, opts ...grpc.CallOption) (*CategoryTreeListResp, error) {
	out := new(CategoryTreeListResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoryTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoryShow(ctx context.Context, in *CategoryShowReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	out := new(CategoryResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoryShow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategorySave(ctx context.Context, in *CategorySaveReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	out := new(CategoryResp)
	err := c.cc.Invoke(ctx, CategoryService_CategorySave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoryDel(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*CategoryResp, error) {
	out := new(CategoryResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoryDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	CategoryGet(context.Context, *CategoryGetReq) (*CategoryListResp, error)
	CategoryTree(context.Context, *CategoryTreeReq) (*CategoryTreeListResp, error)
	CategoryShow(context.Context, *CategoryShowReq) (*CategoryResp, error)
	CategorySave(context.Context, *CategorySaveReq) (*CategoryResp, error)
	CategoryDel(context.Context, *CategoryDelReq) (*CategoryResp, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) CategoryGet(context.Context, *CategoryGetReq) (*CategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryGet not implemented")
}
func (UnimplementedCategoryServiceServer) CategoryTree(context.Context, *CategoryTreeReq) (*CategoryTreeListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryTree not implemented")
}
func (UnimplementedCategoryServiceServer) CategoryShow(context.Context, *CategoryShowReq) (*CategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryShow not implemented")
}
func (UnimplementedCategoryServiceServer) CategorySave(context.Context, *CategorySaveReq) (*CategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategorySave not implemented")
}
func (UnimplementedCategoryServiceServer) CategoryDel(context.Context, *CategoryDelReq) (*CategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryDel not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CategoryGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryGet(ctx, req.(*CategoryGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoryTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryTreeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryTree(ctx, req.(*CategoryTreeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoryShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryShowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryShow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryShow(ctx, req.(*CategoryShowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategorySave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategorySaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategorySave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategorySave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategorySave(ctx, req.(*CategorySaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoryDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryDel(ctx, req.(*CategoryDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategoryGet",
			Handler:    _CategoryService_CategoryGet_Handler,
		},
		{
			MethodName: "CategoryTree",
			Handler:    _CategoryService_CategoryTree_Handler,
		},
		{
			MethodName: "CategoryShow",
			Handler:    _CategoryService_CategoryShow_Handler,
		},
		{
			MethodName: "CategorySave",
			Handler:    _CategoryService_CategorySave_Handler,
		},
		{
			MethodName: "CategoryDel",
			Handler:    _CategoryService_CategoryDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/shop.proto",
}

const (
	ProductService_ProductGet_FullMethodName  = "/shop.ProductService/ProductGet"
	ProductService_ProductShow_FullMethodName = "/shop.ProductService/ProductShow"
	ProductService_ProductSave_FullMethodName = "/shop.ProductService/ProductSave"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	ProductGet(ctx context.Context, in *ProductGetReq, opts ...grpc.CallOption) (*ProductListResp, error)
	ProductShow(ctx context.Context, in *ProductShowReq, opts ...grpc.CallOption) (*ProductResp, error)
	ProductSave(ctx context.Context, in *ProductSaveReq, opts ...grpc.CallOption) (*ProductSaveResp, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) ProductGet(ctx context.Context, in *ProductGetReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	out := new(ProductListResp)
	err := c.cc.Invoke(ctx, ProductService_ProductGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ProductShow(ctx context.Context, in *ProductShowReq, opts ...grpc.CallOption) (*ProductResp, error) {
	out := new(ProductResp)
	err := c.cc.Invoke(ctx, ProductService_ProductShow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ProductSave(ctx context.Context, in *ProductSaveReq, opts ...grpc.CallOption) (*ProductSaveResp, error) {
	out := new(ProductSaveResp)
	err := c.cc.Invoke(ctx, ProductService_ProductSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	ProductGet(context.Context, *ProductGetReq) (*ProductListResp, error)
	ProductShow(context.Context, *ProductShowReq) (*ProductResp, error)
	ProductSave(context.Context, *ProductSaveReq) (*ProductSaveResp, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) ProductGet(context.Context, *ProductGetReq) (*ProductListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductGet not implemented")
}
func (UnimplementedProductServiceServer) ProductShow(context.Context, *ProductShowReq) (*ProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductShow not implemented")
}
func (UnimplementedProductServiceServer) ProductSave(context.Context, *ProductSaveReq) (*ProductSaveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSave not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_ProductGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ProductGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductGet(ctx, req.(*ProductGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ProductShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductShowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ProductShow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductShow(ctx, req.(*ProductShowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ProductSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ProductSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductSave(ctx, req.(*ProductSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductGet",
			Handler:    _ProductService_ProductGet_Handler,
		},
		{
			MethodName: "ProductShow",
			Handler:    _ProductService_ProductShow_Handler,
		},
		{
			MethodName: "ProductSave",
			Handler:    _ProductService_ProductSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/shop.proto",
}

const (
	ProductAttrService_AttrKeySave_FullMethodName = "/shop.ProductAttrService/AttrKeySave"
	ProductAttrService_AttrValSave_FullMethodName = "/shop.ProductAttrService/AttrValSave"
)

// ProductAttrServiceClient is the client API for ProductAttrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductAttrServiceClient interface {
	AttrKeySave(ctx context.Context, in *ProductAttrKeyReq, opts ...grpc.CallOption) (*ProductAttrKeyResp, error)
	AttrValSave(ctx context.Context, in *ProductAttrValReq, opts ...grpc.CallOption) (*ProductAttrValResp, error)
}

type productAttrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductAttrServiceClient(cc grpc.ClientConnInterface) ProductAttrServiceClient {
	return &productAttrServiceClient{cc}
}

func (c *productAttrServiceClient) AttrKeySave(ctx context.Context, in *ProductAttrKeyReq, opts ...grpc.CallOption) (*ProductAttrKeyResp, error) {
	out := new(ProductAttrKeyResp)
	err := c.cc.Invoke(ctx, ProductAttrService_AttrKeySave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productAttrServiceClient) AttrValSave(ctx context.Context, in *ProductAttrValReq, opts ...grpc.CallOption) (*ProductAttrValResp, error) {
	out := new(ProductAttrValResp)
	err := c.cc.Invoke(ctx, ProductAttrService_AttrValSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductAttrServiceServer is the server API for ProductAttrService service.
// All implementations must embed UnimplementedProductAttrServiceServer
// for forward compatibility
type ProductAttrServiceServer interface {
	AttrKeySave(context.Context, *ProductAttrKeyReq) (*ProductAttrKeyResp, error)
	AttrValSave(context.Context, *ProductAttrValReq) (*ProductAttrValResp, error)
	mustEmbedUnimplementedProductAttrServiceServer()
}

// UnimplementedProductAttrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductAttrServiceServer struct {
}

func (UnimplementedProductAttrServiceServer) AttrKeySave(context.Context, *ProductAttrKeyReq) (*ProductAttrKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttrKeySave not implemented")
}
func (UnimplementedProductAttrServiceServer) AttrValSave(context.Context, *ProductAttrValReq) (*ProductAttrValResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttrValSave not implemented")
}
func (UnimplementedProductAttrServiceServer) mustEmbedUnimplementedProductAttrServiceServer() {}

// UnsafeProductAttrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductAttrServiceServer will
// result in compilation errors.
type UnsafeProductAttrServiceServer interface {
	mustEmbedUnimplementedProductAttrServiceServer()
}

func RegisterProductAttrServiceServer(s grpc.ServiceRegistrar, srv ProductAttrServiceServer) {
	s.RegisterService(&ProductAttrService_ServiceDesc, srv)
}

func _ProductAttrService_AttrKeySave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAttrKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAttrServiceServer).AttrKeySave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductAttrService_AttrKeySave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAttrServiceServer).AttrKeySave(ctx, req.(*ProductAttrKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductAttrService_AttrValSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAttrValReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAttrServiceServer).AttrValSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductAttrService_AttrValSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAttrServiceServer).AttrValSave(ctx, req.(*ProductAttrValReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductAttrService_ServiceDesc is the grpc.ServiceDesc for ProductAttrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductAttrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ProductAttrService",
	HandlerType: (*ProductAttrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttrKeySave",
			Handler:    _ProductAttrService_AttrKeySave_Handler,
		},
		{
			MethodName: "AttrValSave",
			Handler:    _ProductAttrService_AttrValSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/shop.proto",
}
