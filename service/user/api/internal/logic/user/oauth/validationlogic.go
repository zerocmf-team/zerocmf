package oauth

import (
	"context"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/user/rpc/types/user"
	"strings"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidationLogic {
	return &ValidationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidationLogic) Validation(req *types.ValidationReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	userRpc := c.UserRpc

	auth := r.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = r.FormValue("access_token")
	}

	if token == "" {
		resp.Error("token不能为空！", nil)
		return
	}

	tenantId := req.TenantId

	result, err := userRpc.ValidationJwt(l.ctx, &user.OauthRequest{
		Token:    token,
		TenantId: tenantId,
	})

	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data.H{
		"userId": result.GetUserId(),
	})
	return
}
