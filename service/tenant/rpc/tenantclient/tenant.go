// Code generated by goctl. DO NOT EDIT.
// Source: tenant.proto

package tenantclient

import (
	"context"

	"zerocmf/service/tenant/rpc/types/tenant"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckUserReply      = tenant.CheckUserReply
	CheckUserReq        = tenant.CheckUserReq
	CurrentUserReq      = tenant.CurrentUserReq
	RegisterReq         = tenant.RegisterReq
	RemoveSiteUserReply = tenant.RemoveSiteUserReply
	RemoveSiteUserReq   = tenant.RemoveSiteUserReq
	UserReply           = tenant.UserReply

	Tenant interface {
		Get(ctx context.Context, in *CurrentUserReq, opts ...grpc.CallOption) (*UserReply, error)
		RegisterUser(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error)
		CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserReply, error)
		RemoveSiteUser(ctx context.Context, in *RemoveSiteUserReq, opts ...grpc.CallOption) (*RemoveSiteUserReply, error)
	}

	defaultTenant struct {
		cli zrpc.Client
	}
)

func NewTenant(cli zrpc.Client) Tenant {
	return &defaultTenant{
		cli: cli,
	}
}

func (m *defaultTenant) Get(ctx context.Context, in *CurrentUserReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := tenant.NewTenantClient(m.cli.Conn())
	return client.Get(ctx, in, opts...)
}

func (m *defaultTenant) RegisterUser(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := tenant.NewTenantClient(m.cli.Conn())
	return client.RegisterUser(ctx, in, opts...)
}

func (m *defaultTenant) CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserReply, error) {
	client := tenant.NewTenantClient(m.cli.Conn())
	return client.CheckUser(ctx, in, opts...)
}

func (m *defaultTenant) RemoveSiteUser(ctx context.Context, in *RemoveSiteUserReq, opts ...grpc.CallOption) (*RemoveSiteUserReply, error) {
	client := tenant.NewTenantClient(m.cli.Conn())
	return client.RemoveSiteUser(ctx, in, opts...)
}
