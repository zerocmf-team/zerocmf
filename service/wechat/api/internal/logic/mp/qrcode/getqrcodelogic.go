package qrcode

import (
	"context"
	"github.com/zerocmf/wechat-easy-sdk/mp/promote"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"strconv"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
	weUtil "zerocmf/service/wechat/api/util"
)

type GetQrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQrcodeLogic {
	return &GetQrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQrcodeLogic) GetQrcode() (resp types.Response) {

	c := l.svcCtx
	appId := c.AppId
	secret := c.Secret

	tokenI, exist := c.Get("token")
	if exist == false {
		resp.Error("系统错误。请稍后再试。", tokenI)
		return
	}

	token, _ := tokenI.(string)

	r := c.Request
	w := c.ResponseWriter
	store := c.Store
	session, _ := store.Get(r, "qrcode")
	values := session.Values

	if values["ticket"] == nil {
		redis := c.Redis
		key := "mp_qrcode"
		reply, err := redis.Incr(key).Result()
		if err != nil {
			resp.Error("获取失败", err.Error())
			return
		}

		expireSeconds := 120

		redis.Expire(key, time.Duration(expireSeconds)*time.Second)

		unix := time.Now().Unix()
		unixStr := strconv.FormatInt(unix, 10)
		replyStr := strconv.FormatInt(reply, 10)

		// 创建唯一的二维码标识
		qrcode := unixStr + replyStr
		qrcodeMd5 := util.GetMd5(qrcode)
		// 调用接口生成二维码
		mpQrcode := new(promote.Qrcode)
		res, err := mpQrcode.Create(token, expireSeconds, mpQrcode.QrStrScene(), mpQrcode.WithSceneStr(qrcode))

		if err != nil {
			resp.Error("系统请求错误。请联系管理员或稍后再试！", err.Error())
			return
		}

		if res.ErrCode == -1 {
			resp.Error(res.ErrMsg, nil)
			return
		}

		if res.ErrCode != 0 {
			session.Options.MaxAge = 0 //失效时间
			weUtil.MpToken(redis, appId, secret, true)
		} else {
			ticket := res.Ticket
			expSec := res.ExpireSeconds
			expAt := time.Now().Unix() + int64(expSec)

			session.Values["qrcode"] = qrcode
			session.Values["ticket"] = ticket
			session.Values["expSec"] = expSec
			session.Values["expAt"] = expAt
			session.Values[qrcodeMd5] = qrcode
			session.Options.MaxAge = expSec //失效时间
		}

		session.Save(r, w)
	}

	ticket := values["ticket"]
	ticketStr := ""

	if ticket != nil {
		var ok bool
		ticketStr, ok = ticket.(string)
		if ok == false {
			resp.Error("系统错误，获取ticket失败", nil)
			return
		}
	}

	if ticketStr != "" {
		ticketStr = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticketStr)
	}

	expSec := values["expSec"]
	expAt := values["expAt"]

	resp.Success("获取成功！", map[string]interface{}{"qrcode": ticketStr, "expire_seconds": expSec, "expire_at": expAt})
	return
}
