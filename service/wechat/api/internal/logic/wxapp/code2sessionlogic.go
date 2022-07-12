package wxapp

import (
	"context"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
	"github.com/daifuyang/wechat-easysdk-go/sns"

	"github.com/zeromicro/go-zero/core/logx"
)

type Code2SessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCode2SessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Code2SessionLogic {
	return &Code2SessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Code2SessionLogic) Code2Session(req *types.Code2SessionReq) (resp types.Response) {
	code := req.JsCode
	if code == "" {
		resp.Error("jscode不能为空！", nil)
		return
	}

	appId := ""
	secret := ""

	code2Session := sns.Code2Session{
		AppId:  appId,
		Secret: secret,
		JsCode: code,
	}

	res, err := sns.Code2session(code2Session)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", res)
	return
}
