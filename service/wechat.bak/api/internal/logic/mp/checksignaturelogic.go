package mp

import (
	"context"
	"sort"
	"strings"
	"zerocmf/common/bootstrap/util"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSignatureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSignatureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSignatureLogic {
	return &CheckSignatureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSignatureLogic) CheckSignature(req *types.CheckSignatureReq) (resp types.Response) {
	signature := req.Signature
	if signature == "" {
		resp.Error("signature不能为空", nil)
		return
	}
	timestamp := req.Timestamp
	if timestamp == "" {
		resp.Error("timestamp不能为空", nil)
		return
	}
	nonce := req.Nonce
	if nonce == "" {
		resp.Error("nonce不能为空", nil)
		return
	}
	token := "gincmf2021"
	echoStr := req.Echostr
	if echoStr == "" {
		resp.Error("echostr不能为空", nil)
		return
	}
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := strings.Join(tmpArr, "")
	tmpSha1 := util.GetSha1(tmpStr)

	if signature == tmpSha1 {
		resp.Success("校验成功", echoStr)
		return
	}
	resp.Error("校验失败", nil)
	return
}
