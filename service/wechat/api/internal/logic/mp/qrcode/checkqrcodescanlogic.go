package qrcode

import (
	"context"
	"strings"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckQrcodeScanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckQrcodeScanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckQrcodeScanLogic {
	return &CheckQrcodeScanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckQrcodeScanLogic) CheckQrcodeScan(req *types.CheckQrcodeScanReq) (resp types.Response) {

	c := l.svcCtx
	r := c.Request
	redis := c.Redis
	store := c.Store

	sceneId := req.SceneId

	if len(strings.TrimSpace(sceneId)) == 0 {
		resp.Error("sceneId不能为空！", nil)
		return
	}

	session, _ := store.Get(r, "qrcode")
	qrcode := session.Values[sceneId]

	if qrcode == nil {
		resp.Error("二维码已过期", nil)
		return
	}

	cmd := redis.Get(qrcode.(string))
	data := cmd.Val()

	resp.Success("获取成功！", data)

	return
}
