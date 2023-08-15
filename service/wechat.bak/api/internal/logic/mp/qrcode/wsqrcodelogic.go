package qrcode

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"zerocmf/service/wechat/api/internal/types"
	weUtil "zerocmf/service/wechat/api/util"

	"github.com/zerocmf/wechatEasySdk/mp/promote"

	"zerocmf/service/wechat/api/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type WsQrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func NewWsQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsQrcodeLogic {
	return &WsQrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsQrcodeLogic) WsQrcode(w http.ResponseWriter, r *http.Request) {
	c := l.svcCtx
	redis := c.Redis

	var (
		conn *websocket.Conn
		err  error
		// msgType int
		//msg []byte
		reply int64
		//bs    []byte
		res     promote.CreateResponse
		marshal []byte
	)

	expireSeconds := 120

	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	values := make(map[string]interface{}, 0)
	for {
		var token string
		appId := c.AppId
		secret := c.Secret

		token, err = weUtil.MpToken(redis, appId, secret, false)
		resp := types.Response{}

		if values["ticket"] == nil {

			key := "mp_qrcode"
			reply, err = redis.Incr(key).Result()
			if err != nil {
				fmt.Println("redis reply err", err.Error())
				return
			}

			redis.Expire(key, time.Duration(expireSeconds)*time.Second)

			unix := time.Now().Unix()
			unixStr := strconv.FormatInt(unix, 10)
			replyStr := strconv.FormatInt(reply, 10)

			// 创建唯一的二维码标识
			qrcode := unixStr + replyStr
			//qrcodeMd5 := util.GetMd5(qrcode)
			// 调用接口生成二维码
			mpQrcode := new(promote.Qrcode)
			res, err = mpQrcode.Create(token, expireSeconds, mpQrcode.QrStrScene(), mpQrcode.WithSceneStr(qrcode))

			if err != nil {
				fmt.Println("mpQrcode err", err.Error())
				return
			}

			if res.ErrCode == -1 {
				return
			}

			if res.ErrCode != 0 {
				values = make(map[string]interface{}, 0)
				weUtil.MpToken(redis, appId, secret, true)
				continue
			} else {
				ticket := res.Ticket
				expSec := expireSeconds
				expAt := time.Now().Unix() + int64(expSec)
				values["qrcode"] = qrcode
				values["ticket"] = ticket
				values["expSec"] = expSec
				values["expAt"] = expAt
			}
		} else {
			unix := time.Now().Unix()
			expAt := values["expAt"]
			expAtInt := expAt.(int64)
			if expAtInt <= unix {
				values = make(map[string]interface{}, 0)
				conn.Close()
				break
			}
		}

		qrcode := values["qrcode"].(string)
		mapVal := redis.HGetAll(qrcode).Val()
		fmt.Println("mapVal", mapVal)
		if len(mapVal) > 0 {
			resp.Success("获取成功！", mapVal)
		} else {
			ticket := values["ticket"].(string)
			qrcodeUrl := "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket)
			resp.Success("获取成功！", map[string]interface{}{"qrcode": qrcodeUrl, "expire_seconds": values["expSec"], "expire_at": values["expAt"]})
		}
		marshal, err = json.Marshal(resp)
		if err != nil {
			return
		}
		if err = conn.WriteMessage(websocket.TextMessage, marshal); err != nil {
			fmt.Println("err", err.Error())
			return
		}
		time.Sleep(time.Second * 1)
	}
}
