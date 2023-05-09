// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: admin.proto

package admin

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
	Admin_GetMenus_FullMethodName    = "/admin.admin/getMenus"
	Admin_AutoMigrate_FullMethodName = "/admin.admin/autoMigrate"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	GetMenus(ctx context.Context, in *AdminMenuReq, opts ...grpc.CallOption) (*AdminMenuReply, error)
	AutoMigrate(ctx context.Context, in *SiteReq, opts ...grpc.CallOption) (*SiteReply, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetMenus(ctx context.Context, in *AdminMenuReq, opts ...grpc.CallOption) (*AdminMenuReply, error) {
	out := new(AdminMenuReply)
	err := c.cc.Invoke(ctx, Admin_GetMenus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AutoMigrate(ctx context.Context, in *SiteReq, opts ...grpc.CallOption) (*SiteReply, error) {
	out := new(SiteReply)
	err := c.cc.Invoke(ctx, Admin_AutoMigrate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	GetMenus(context.Context, *AdminMenuReq) (*AdminMenuReply, error)
	AutoMigrate(context.Context, *SiteReq) (*SiteReply, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) GetMenus(context.Context, *AdminMenuReq) (*AdminMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenus not implemented")
}
func (UnimplementedAdminServer) AutoMigrate(context.Context, *SiteReq) (*SiteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutoMigrate not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_GetMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetMenus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMenus(ctx, req.(*AdminMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AutoMigrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AutoMigrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AutoMigrate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AutoMigrate(ctx, req.(*SiteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMenus",
			Handler:    _Admin_GetMenus_Handler,
		},
		{
			MethodName: "autoMigrate",
			Handler:    _Admin_AutoMigrate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
