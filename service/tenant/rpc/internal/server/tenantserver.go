// Code generated by goctl. DO NOT EDIT.
// Source: tenant.proto

package server

import (
	"context"

	"zerocmf/service/tenant/rpc/internal/logic"
	"zerocmf/service/tenant/rpc/internal/svc"
	"zerocmf/service/tenant/rpc/types/tenant"
)

type TenantServer struct {
	svcCtx *svc.ServiceContext
	tenant.UnimplementedTenantServer
}

func NewTenantServer(svcCtx *svc.ServiceContext) *TenantServer {
	return &TenantServer{
		svcCtx: svcCtx,
	}
}

func (s *TenantServer) Get(ctx context.Context, in *tenant.CurrentUserReq) (*tenant.UserReply, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}
