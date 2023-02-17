// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"zerocmf/service/user/rpc/internal/logic"
	"zerocmf/service/user/rpc/internal/svc"
	"zerocmf/service/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Get(ctx context.Context, in *user.UserRequest) (*user.UserReply, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *UserServer) Init(ctx context.Context, in *user.InitRequest) (*user.InitReply, error) {
	l := logic.NewInitLogic(ctx, s.svcCtx)
	return l.Init(in)
}

func (s *UserServer) ValidationJwt(ctx context.Context, in *user.OauthRequest) (*user.OauthReply, error) {
	l := logic.NewValidationJwtLogic(ctx, s.svcCtx)
	return l.ValidationJwt(in)
}

func (s *UserServer) Database(ctx context.Context, in *user.DatabaseRequest) (*user.DatabaseReply, error) {
	l := logic.NewDatabaseLogic(ctx, s.svcCtx)
	return l.Database(in)
}
