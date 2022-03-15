// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package oauth

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

// OauthClient is the client API for Oauth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OauthClient interface {
	// Sends a greeting
	ValidationBearerToken(ctx context.Context, in *OauthRequest, opts ...grpc.CallOption) (*OauthReply, error)
}

type oauthClient struct {
	cc grpc.ClientConnInterface
}

func NewOauthClient(cc grpc.ClientConnInterface) OauthClient {
	return &oauthClient{cc}
}

func (c *oauthClient) ValidationBearerToken(ctx context.Context, in *OauthRequest, opts ...grpc.CallOption) (*OauthReply, error) {
	out := new(OauthReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/ValidationBearerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthServer is the server API for Oauth service.
// All implementations must embed UnimplementedOauthServer
// for forward compatibility
type OauthServer interface {
	// Sends a greeting
	ValidationBearerToken(context.Context, *OauthRequest) (*OauthReply, error)
	mustEmbedUnimplementedOauthServer()
}

// UnimplementedOauthServer must be embedded to have forward compatible implementations.
type UnimplementedOauthServer struct {
}

func (UnimplementedOauthServer) ValidationBearerToken(context.Context, *OauthRequest) (*OauthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidationBearerToken not implemented")
}
func (UnimplementedOauthServer) mustEmbedUnimplementedOauthServer() {}

// UnsafeOauthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OauthServer will
// result in compilation errors.
type UnsafeOauthServer interface {
	mustEmbedUnimplementedOauthServer()
}

func RegisterOauthServer(s grpc.ServiceRegistrar, srv OauthServer) {
	s.RegisterService(&Oauth_ServiceDesc, srv)
}

func _Oauth_ValidationBearerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OauthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).ValidationBearerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/ValidationBearerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).ValidationBearerToken(ctx, req.(*OauthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth_ServiceDesc is the grpc.ServiceDesc for Oauth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.Oauth",
	HandlerType: (*OauthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidationBearerToken",
			Handler:    _Oauth_ValidationBearerToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}
