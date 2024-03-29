// Code generated by goctl. DO NOT EDIT.
// Source: lowcode.proto

package server

import (
	"context"

	"zerocmf/service/lowcode/rpc/internal/logic"
	"zerocmf/service/lowcode/rpc/internal/svc"
	"zerocmf/service/lowcode/rpc/types/lowcode"
)

type LowcodeServer struct {
	svcCtx *svc.ServiceContext
	lowcode.UnimplementedLowcodeServer
}

func NewLowcodeServer(svcCtx *svc.ServiceContext) *LowcodeServer {
	return &LowcodeServer{
		svcCtx: svcCtx,
	}
}

func (s *LowcodeServer) AutoMigrate(ctx context.Context, in *lowcode.SiteReq) (*lowcode.SiteReply, error) {
	l := logic.NewAutoMigrateLogic(ctx, s.svcCtx)
	return l.AutoMigrate(in)
}
